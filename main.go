package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{
		fileserverHits: 0,
	}

	//mux := http.NewServeMux()
	r := chi.NewRouter()
	r.Mount("/", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(filepathRoot))))
	r.Get("/healthz", handlerReadiness)
	r.Get("/metrics", apiCfg.handlerMetrics)

	corsMux := middlewareCors(r)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
