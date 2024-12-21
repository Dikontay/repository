package http

import (
	"net/http"
	"repository/internal/endpoints/http/create"
	"repository/internal/endpoints/http/delete"
	"repository/internal/endpoints/http/read"
	"repository/internal/endpoints/http/update"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", create.Handler())
	mux.HandleFunc("/read/", read.Handler())
	mux.HandleFunc("/update/", update.Handler())
	mux.HandleFunc("/delete/", delete.Handler())
	return mux
}
