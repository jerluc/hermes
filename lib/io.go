package hermes

import (
	// "bytes"
	"io"
)

func NewCopyBuffer(writers...io.Writer) io.Writer {
	return writers[0]
}