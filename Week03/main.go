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

func main() {
	done := make(chan int)
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		handler := http.NewServeMux()
		handler.HandleFunc("/close", func(writer http.ResponseWriter, request *http.Request) {
			done <- 1
		})
		s := &http.Server{
			Addr:    ":8080",
			Handler: handler,
		}

		httpEg, httpCtx := errgroup.WithContext(context.Background())
		httpEg.Go(func() error {
			return s.ListenAndServe()
		})
		httpEg.Go(func() error {
			select {
				case <-done:
					fmt.Println("协程1：访问/close路由，触发信号关闭1")
				case <-ctx.Done():
					fmt.Println("协程1：来自信号关闭2")
				case <-httpCtx.Done():
					fmt.Println("http server error")
					return errors.New("http server error")
			}
			timeoutCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)
			return s.Shutdown(timeoutCtx)
		})
		return httpEg.Wait()
	})
	eg.Go(func() error {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-quit:
			fmt.Println("协程2：监听系统信号量，触发信号关闭2")
			return errors.New("协程2：信号关闭2")
		case <-ctx.Done():
			fmt.Println("协程2：来自信号关闭1")
			return errors.New("协程2：【通知】http服务关闭")
		}
	})
	fmt.Println("开始捕捉err")
	err := eg.Wait()
	fmt.Println("=======", err)
}