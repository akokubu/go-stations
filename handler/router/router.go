package router

import (
	"database/sql"
	"net/http"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	// healthzHandler := handler.NewHealthzHandler()
	// mux.HandleFunc(healthzHandler.Endpoint, healthzHandler.ServeHTTP)
	return mux
}
