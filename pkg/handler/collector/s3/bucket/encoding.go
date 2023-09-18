package bucket

import (
	"github.com/guacsec/guac/pkg/handler/processor"
	"strings"
)

const (
	BZIP2   = "BZIP2"
	ZSTD    = "ZSTD"
	UNKNOWN = "UNKNOWN"
)

func ExtractEncoding(encoding string, filename string) processor.EncodingType {
	switch encoding {
	case BZIP2:
		return BZIP2
	case ZSTD:
		return ZSTD
	default:
		return FromFile(filename)
	}
}

func FromFile(file string) processor.EncodingType {
	strs := strings.Split(file, ".")
	extension := strs[len(strs)-1]
	switch extension {
	case "bz2":
		return BZIP2
	case "zst":
		return ZSTD
	default:
		return UNKNOWN
	}
}
