package s3

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/guacsec/guac/pkg/handler/collector"
	"github.com/guacsec/guac/pkg/handler/collector/s3/bucket"
	"github.com/guacsec/guac/pkg/handler/collector/s3/messaging"
	"github.com/guacsec/guac/pkg/handler/processor"
	"io"
	"os"
	"os/signal"
	"strings"
	"testing"
	"time"
)

// Test message
type TestMessage struct {
	item   string
	bucket string
	event  messaging.EventName
}

func (msg TestMessage) GetEvent() (messaging.EventName, error) {
	return msg.event, nil
}

func (msg TestMessage) GetBucket() (string, error) {
	return msg.bucket, nil
}

func (msg TestMessage) GetItem() (string, error) {
	return msg.item, nil
}

// Test Message Provider
type TestProvider struct {
	queue string
}

func NewTestProvider(queue string) TestProvider {
	return TestProvider{queue}
}

func (t TestProvider) ReceiveMessage(context.Context) (messaging.Message, error) {
	fmt.Printf("returning message for queue %s\n", t.queue)
	time.Sleep(2 * time.Second)

	return TestMessage{
		item:   "item",
		bucket: t.queue,
		event:  messaging.PUT,
	}, nil
}

func (t TestProvider) Close(ctx context.Context) error {
	return nil
}

// Test Message Provider builder
type TestMpBuilder struct {
}

func (tb TestMpBuilder) GetMessageProvider(config messaging.MessageProviderConfig) (messaging.MessageProvider, error) {
	return NewTestProvider(config.Queue), nil
}

// Test Bucket
type TestBucket struct {
}

func (td TestBucket) DownloadFile(ctx context.Context, bucket string, item string) ([]byte, error) {
	fmt.Printf("downloading file with bucket %s and name %s\n", bucket, item)
	return []byte{1, 2, 3}, nil
}

func (td TestBucket) GetEncoding(ctx context.Context, bucket string, item string) (string, error) {
	return "", nil
}

type TestBucketBuilder struct {
}

func (td TestBucketBuilder) GetDownloader(hostname string, port string, region string) bucket.Bucket {
	return TestBucket{}
}

func TestQueuesSplit(t *testing.T) {
	ctx := context.Background()

	sigChan := make(chan os.Signal)
	s3Collector, _ := NewS3Collector(S3CollectorConfig{
		Queues:        "q1,q2",
		MpBuilder:     TestMpBuilder{},
		BucketBuilder: TestBucketBuilder{},
		SigChan:       sigChan,
	})

	if err := collector.RegisterDocumentCollector(s3Collector, S3CollectorType); err != nil &&
		!errors.Is(err, collector.ErrCollectorOverwrite) {
		t.Fatalf("could not register collector: %v", err)
	}

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	var s []*processor.Document
	em := func(d *processor.Document) error {
		s = append(s, d)
		return nil
	}
	eh := func(err error) bool {
		return true
	}

	go collector.Collect(ctx, em, eh)
	time.Sleep(5 * time.Second)
	signal.Notify(sigChan, os.Interrupt)

	w.Close()
	os.Stdout = oldStdout // restoring the real stdout
	out := <-outC

	fmt.Println(out)

	if !strings.Contains(out, "returning message for queue q1") {
		t.Errorf("message for q1 not returned")
	}
	if !strings.Contains(out, "returning message for queue q2") {
		t.Errorf("message for q2 not returned")
	}
	if !strings.Contains(out, "downloading file with bucket q1 and name item") {
		t.Errorf("not downloading from bucket q1")
	}
	if !strings.Contains(out, "downloading file with bucket q2 and name item") {
		t.Errorf("not downloading from bucket q2")
	}
}
