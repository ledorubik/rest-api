package middlewares

import (
	"bytes"
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var buf bytes.Buffer
		tee := io.TeeReader(ctx.Request.Body, &buf)
		body, _ := io.ReadAll(tee)
		ctx.Request.Body = io.NopCloser(&buf)
		log.Printf("request URL: %s", ctx.Request.URL)
		log.Printf("request headers: %s", ctx.Request.Header)
		log.Printf("request body: %s", strings.ReplaceAll(string(body), "\n", ""))
		ctx.Next()
	}
}
