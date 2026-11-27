// Package httpapi doc here
package httpapi

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_, _ = w.Write([]byte(`{"status":"OK"}`))
	})

	return mux
}
