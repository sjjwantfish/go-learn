package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Service struct {
	Name string
	Addr string
	http.Server
}

func (s *Service) StartListen() error {
	s.Server.Addr = s.Addr
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	return s.Server.ListenAndServe()
}

func main() {
	// 基于 errgroup 实现一个 http server 的启动和关闭 ，
	// 以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。
	g, ctx := errgroup.WithContext(context.Background())

	app := &Service{
		Name: "App",
		Addr: ":8080",
	}

	g.Go(func() error {
		log.Printf("Startting %v serve...", app.Name)
		return app.StartListen()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		}
		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		log.Println("Stopping all serve...")
		return app.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}
