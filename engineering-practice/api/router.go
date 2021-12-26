package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"time"
)

func Success(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, gin.H{"retcode": 0, "resp": resp, "request": c.Request.URL.Path})
}

func RecoverErr() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				const size = 4096
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]

				logrus.WithFields(logrus.Fields{
					"url": c.Request.URL.Path,
				}).Error("发生错误:", err, "  方法: ", c.HandlerName(), " stack: ", string(buf))
				c.AbortWithStatusJSON(500, gin.H{"error": "sorry, we made a mistake!", "retcode": 1001, "url": c.Request.URL.Path})
			}
		}()
		c.Next()
	}
}

func LoadRouter(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	g.Use(RecoverErr())
	g.RedirectTrailingSlash = false
	g.HandleMethodNotAllowed = true
	g.NoMethod(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "method not allowed", "retcode": 1005, "request": c.Request.URL.Path})
	})
	g.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "route not found", "retcode": 1007, "request": c.Request.URL.Path})
	})
	g.GET("/healthz", func(c *gin.Context) {
		Success(c, gin.H{"status": "ok"})
		return
	})
	g.GET("/_count", func(c *gin.Context) {
		Success(c, gin.H{"goroutine count": runtime.NumGoroutine()})
		return
	})

	svcd := g.Group("/api/v1/")
	{
		svcd.POST("/user/create", CreateUser)
	}


	return g
}