package messaging

import "context"

type EventName string

const (
	PUT EventName = "PUT"
)

// Message A generic message related to an S3 bucket and item
type Message interface {
	GetEvent() (EventName, error)
	GetBucket() (string, error)
	GetItem() (string, error)
}

// MessageProvider Reads and returns messages from a given queue (or topic)
type MessageProvider interface {
	ReceiveMessage(ctx context.Context) (Message, error)
	Close(ctx context.Context) error
}

type MessageProviderConfig struct {
	Queue    string
	Host     string
	Port     string
	Provider string
	Region   string
}

// MessageProviderBuilder Returns a builder for a MessageProvider
type MessageProviderBuilder interface {
	GetMessageProvider(config MessageProviderConfig) (MessageProvider, error)
}

type MpBuilder struct {
}

// GetMessageProvider Returns a MessageProvider with the given config. Defaults to Kafka provider if no MESSAGE_PROVIDER environment variable is found
func (mb MpBuilder) GetMessageProvider(config MessageProviderConfig) (MessageProvider, error) {
	switch config.Provider {
	case "sqs":
		return NewSqsProvider(config)
	default:
		return NewKafkaProvider(config)
	}
}

func GetDefaultMessageProviderBuilder() (MessageProviderBuilder, error) {
	return MpBuilder{}, nil
}
