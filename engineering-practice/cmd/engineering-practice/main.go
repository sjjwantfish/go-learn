package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sjjwantfish/go-learn/engineering-practice/api"
	"github.com/sjjwantfish/go-learn/engineering-practice/internal/data"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Message struct {
	msg string
}
type Greeter struct {
	Message Message
}
type Event struct {
	Greeter Greeter
}
// NewMessage Message的构造函数
func NewMessage(msg string) Message {
	return Message{
		msg:msg,
	}
}
// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}
// NewEvent Event构造函数
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}
func (g Greeter) Greet() Message {
	return g.Message
}


func main() {
	// 按照自己的构想，写一个项目满足基本的目录结构和工程，
	// 代码需要包含对数据层、业务层、API 注册，以及 main 函数对于服务的注册和启动，
	// 信号处理，使用 Wire 构建依赖。可以使用自己熟悉的框架。

	err := data.InitDB()
	if err != nil {
		log.Fatal("db init failed")
	}
	InitializeEvent("hello, world!")
	g := gin.New()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	api.LoadRouter(g)
	eg, ctx := errgroup.WithContext(context.Background())

	eg.Go(func() error {
		log.Printf("Startting serve...")
		return g.Run()
	})

	eg.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		log.Println("Stopping all serve...")
		return server.Shutdown(timeoutCtx)
	})

	eg.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("get os signal: %v", sig)
		}
	})


}
