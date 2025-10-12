package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/rrajvardhan/bunkr/internal/tui/shared"
)

var srv *http.Server

func Start(port string, sState *shared.ServerState) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		sState.URL = "http://localhost" + port
		sState.Err = fmt.Sprintf("Could not get local IP: %v", err)
	} else {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				sState.URL = fmt.Sprintf("http://%s%s", ipnet.IP.String(), port)
				break
			}
		}
		if sState.URL == "" {
			sState.URL = "http://localhost" + port
		}
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/ping", PingHandler)
	mux.HandleFunc("/upload", UploadHandler)

	srv = &http.Server{
		Addr:    "0.0.0.0" + port,
		Handler: mux,
	}

	go func() {
		sState.Running = true
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sState.Running = false
			sState.Err = fmt.Sprintf("Server error: %v", err)
		}
	}()
}

func Quit() {
	if srv != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_ = srv.Shutdown(ctx)
	}

	if err := os.RemoveAll("./uploads"); err != nil {
		fmt.Printf("failed to remove uploads folder: %v\n", err)
	}
}
