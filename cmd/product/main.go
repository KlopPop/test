package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"product/internal/config"
	"product/internal/generator"
	"product/internal/product"
	"product/internal/sl"
	"product/internal/storage"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.LogFormat, cfg.LogLevel)

	log.Info("Starting app: 'Product'", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	s, err := storage.New(log)

	if err != nil {
		log.Error("Failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	defer s.Close()

	generator.Generate(log, s)

	log.Info("Starting server", slog.String("address", cfg.HTTPServer.Address))
	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(time.Duration(cfg.Delay) * time.Second)
		//log.Info("resp ", slog.String("time", time.Now().Format(time.DateTime)))
		w.WriteHeader(200)
		w.Write([]byte("hello world"))
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	router.Get("/products", func(w http.ResponseWriter, r *http.Request) {
		log.Debug("GET products", "params were:", r.URL.Query())

		// if only one expected
		id, err := strconv.Atoi(r.URL.Query().Get("id"))

		if (err != nil) && (id == 0) {
			w.Write([]byte("Wrong request"))
		} else {
			pr, err := product.GetProduct(id, log, s)

			if err != nil {
				w.Write([]byte("No data found"))
			} else {
				ret := fmt.Sprint("Product id", pr.Id, "\nPrice:", pr.Price, "\n")

				log.Debug("GetProduct", "pr.Id", pr.Id)
				for _, attr := range pr.Attrs {
					log.Debug("GetProduct", attr.Key, attr.Value)
					ret = fmt.Sprint(ret, attr.Key, ": ", attr.Value, "\n")
				}
				w.Write([]byte(ret))
			}

		}

	})

	log.Info("Server started")

	err = http.ListenAndServe(cfg.HTTPServer.Address, router)

	if err != nil {
		log.Error("Failed to listen "+cfg.HTTPServer.Address, sl.Err(err))
		os.Exit(1)
	}

	log.Info("Server stopped")

}