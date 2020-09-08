package example

import (
	"fmt"
	"net/http"
	"time"

	brotli "github.com/freyo/gin-cbrotli"
	"github.com/gin-gonic/gin"
	"github.com/google/brotli/go/cbrotli"
)

func main() {
	r := gin.Default()
	r.Use(brotli.Brotli(cbrotli.WriterOptions{Quality: 5}))
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
	})

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
