package brotli

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/brotli/go/cbrotli"
)

type brotliHandler struct {
	*Options
	brPool sync.Pool
}

func newBrotliHandler(brOption cbrotli.WriterOptions, options ...Option) *brotliHandler {
	var brPool sync.Pool
	brPool.New = func() interface{} {
		br := cbrotli.NewWriter(ioutil.Discard, brOption)
		return &cbrotliWriter{br, brOption}
	}
	handler := &brotliHandler{
		Options: DefaultOptions,
		brPool:  brPool,
	}
	for _, setter := range options {
		setter(handler.Options)
	}
	return handler
}

func (br *brotliHandler) Handle(c *gin.Context) {
	if fn := br.DecompressFn; fn != nil && c.Request.Header.Get("Content-Encoding") == "br" {
		fn(c)
	}

	if len(c.Writer.Header().Get("Content-Encoding")) > 0 {
		return
	}

	if !br.shouldCompress(c.Request) {
		return
	}

	brWriter := br.brPool.Get().(*cbrotliWriter)
	defer br.brPool.Put(brWriter)
	defer brWriter.Reset(ioutil.Discard)
	brWriter.Reset(c.Writer)

	c.Header("Content-Encoding", "br")
	c.Header("Vary", "Accept-Encoding")
	c.Writer = &brotliWriter{c.Writer, brWriter}
	defer func() {
		brWriter.Close()
		c.Header("Content-Length", fmt.Sprint(c.Writer.Size()))
	}()
	c.Next()
}

func (br *brotliHandler) shouldCompress(req *http.Request) bool {
	if !strings.Contains(req.Header.Get("Accept-Encoding"), "br") ||
		strings.Contains(req.Header.Get("Connection"), "Upgrade") ||
		strings.Contains(req.Header.Get("Content-Type"), "text/event-stream") {

		return false
	}

	extension := filepath.Ext(req.URL.Path)
	if br.ExcludedExtensions.Contains(extension) {
		return false
	}

	if br.ExcludedPaths.Contains(req.URL.Path) {
		return false
	}
	if br.ExcludedPathesRegexs.Contains(req.URL.Path) {
		return false
	}

	return true
}
