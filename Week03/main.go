package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

var done = make(chan int)

func main() {
	ctxContext, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctxContext)
	group.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
			done <- 1
		})
		s := NewServer(":8080", ctx, mux)
		go func() {
			err := s.Start()
			if err != nil {
				fmt.Println(err)
			}
		}()
		select {
		case <-done:
			cancel()
			fmt.Println("协程1：访问/close路由，触发信号关闭1")
			return s.Stop()
		case <-ctx.Done():
			fmt.Println("协程1：来自信号关闭2")
			return errors.New("协程1：【通知】信号关闭")
		}
	})
	group.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			cancel()
			fmt.Println("协程2：监听系统信号量，触发信号关闭2")
			return errors.New("协程2：信号关闭2")
		case <-ctx.Done():
			fmt.Println("协程2：来自信号关闭1")
			return errors.New("协程2：【通知】http服务关闭")
		}
	})

	fmt.Println("开始捕捉err")
	err := group.Wait()
	fmt.Println("=======", err)
}

//http服务
type httpServer struct {
	s *http.Server
	handler http.Handler
	cxt context.Context
}

func NewServer(address string, ctxContext context.Context, mux http.Handler) *httpServer {
	h := &httpServer{cxt: ctxContext}
	h.s = &http.Server{
		Addr: address,
		WriteTimeout: time.Second * 3,
		Handler: mux,
	}
	return h
}

func (h *httpServer) Start() error {
	fmt.Println("httpServer start")
	return h.s.ListenAndServe()
}

func (h *httpServer) Stop() error {
	_ = h.s.Shutdown(h.cxt)
	fmt.Println("httpServer结束1")
	return fmt.Errorf("httpServer结束2")
}