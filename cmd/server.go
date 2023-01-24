package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type Server struct {
	srv *http.Server
	l   net.Listener
}

func NewServer(l net.Listener, router http.Handler) *Server {
	return &Server{
		srv: &http.Server{Handler: router},
		l:   l,
	}
}

func (s *Server) Run(ctx context.Context) error {
	// 割り込みシグナルまたは終了シグナルの受信のために待機する
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		// シャットダウンの実行結果を判定する
		// http.Server.Shutdown()が正常終了した場合のhttp.ErrServerClosedステータスは除外する
		if err := s.srv.Serve(s.l); err != nil && err != http.ErrServerClosed {
			return fmt.Errorf("failed to close: %+v", err)
		}
		return nil
	})

	<-ctx.Done()
	// グレースフルシャットダウンを実行する
	if err := s.srv.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: +%v", err)
	}

	// シャットダウンの実行結果を受ける
	return eg.Wait()
}
