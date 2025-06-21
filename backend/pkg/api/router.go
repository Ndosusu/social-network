package api

import "net/http"

func InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Register the default handlers
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})

	return mux
}
