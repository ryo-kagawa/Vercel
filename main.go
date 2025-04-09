package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/ryo-kagawa/Vercel/api"
	"github.com/ryo-kagawa/Vercel/api/karaoke"
)

const addr = ":8080"

func main() {
	http.HandleFunc("/", api.Handler)
	http.HandleFunc("/karaoke/", karaoke.Handler)
	slog.Info("サーバー起動", slog.String("addr", addr))
	if err := http.ListenAndServe(addr, nil); err != nil {
		slog.Error("エラー", slog.Any("error", err))
		os.Exit(1)
	}
}

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
}
