package brotli

import (
	"io"

	"github.com/google/brotli/go/cbrotli"
)

type cbrotliWriter struct {
	*cbrotli.Writer
	cbrotli.WriterOptions
}

// Reset discards the Writer z's state and makes it equivalent to the
// result of its original state from NewWriter or NewWriterLevel, but
// writing to w instead. This permits reusing a Writer rather than
// allocating a new one.
func (z *cbrotliWriter) Reset(w io.Writer) {
	z.init(w, z.WriterOptions)
}

func (z *cbrotliWriter) init(w io.Writer, opt cbrotli.WriterOptions) {
	*z = cbrotliWriter{
		Writer:        cbrotli.NewWriter(w, opt),
		WriterOptions: opt,
	}
}
