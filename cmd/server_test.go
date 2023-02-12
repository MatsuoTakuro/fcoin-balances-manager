package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestServer_Run(t *testing.T) {
	l, err := net.Listen("tcp", "localhost:18080")
	if err != nil {
		t.Fatalf("failed to listen port %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	r := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "test message: %s", r.URL.Path[1:])
	})

	eg.Go(func() error {
		s := NewServer(l, r)
		// テストサーバを起動する
		return s.Run(ctx)
	})

	path := "ok"
	url := fmt.Sprintf("http://%s/%s", l.Addr().String(), path)
	t.Logf("try to request to %q", url)
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("failed to get: %+v", err)
	}
	defer resp.Body.Close()

	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("failed to read body: %+v", err)
	}

	cancel()
	// テストサーバがグレースフルシャットダウンを実行できたかを検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}

	want := fmt.Sprintf("test message: %s", path)
	if string(got) != want {
		t.Errorf("want %q, but got %q", want, got)
	}
}
