# gin-cbrotli
Google Brotli Gin Middleware

## Usage

Download and install it:

```sh
go get github.com/freyo/gin-cbrotli
```

Import it in your code:

```go
import "github.com/freyo/gin-cbrotli"
```

Canonical example:

```go
package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/freyo/gin-cbrotli"
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
```
