package brotli

import (
	"github.com/gin-gonic/gin"
	"github.com/google/brotli/go/cbrotli"
)

func Brotli(brOption cbrotli.WriterOptions, options ...Option) gin.HandlerFunc {
	return newBrotliHandler(brOption, options...).Handle
}

type brotliWriter struct {
	gin.ResponseWriter
	writer *cbrotliWriter
}

func (br *brotliWriter) WriteString(s string) (int, error) {
	return br.writer.Write([]byte(s))
}

func (br *brotliWriter) Write(data []byte) (int, error) {
	return br.writer.Write(data)
}

// Fix: https://github.com/mholt/caddy/issues/38
func (br *brotliWriter) WriteHeader(code int) {
	br.Header().Del("Content-Length")
	br.ResponseWriter.WriteHeader(code)
}
