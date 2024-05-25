package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/KlopPop/test/app1go/internal/config"
	"github.com/KlopPop/test/app1go/internal/sl"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	cfg := config.MustLoad()
	log := sl.SetupLogger(sl.EnvLocal)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(cfg.Delay) * time.Second)
		//log.Info("resp ", slog.String("time", time.Now().Format(time.DateTime)))
		w.WriteHeader(200)
		w.Write([]byte("hello world"))

	})

	log.Info("starting server", slog.String("address", cfg.ServerAddress))

	//done := make(chan os.Signal, 1)
	//signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("server started")
	http.ListenAndServe(cfg.ServerAddress, router)

	log.Info("stopping server")

	log.Info("server stopped")

}
