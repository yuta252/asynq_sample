package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

func main() {
	log.Printf("start app....")
	r := chi.NewRouter()
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitoring",
		RedisConnOpt: asynq.RedisClientOpt{Addr: ":6379"},
	})
	r.Route("/api", func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("status ok"))
			w.WriteHeader(http.StatusOK)
		})
	})
	r.Mount(h.RootPath(), h)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("could not activate server: %v", err)
	}
	log.Printf("end app....")
}
