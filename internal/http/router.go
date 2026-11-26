// Package httpapi doc here
package httpapi

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	return mux
}
