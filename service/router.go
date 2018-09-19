package service

import (
	"github.com/gorilla/mux"
)

// NewRouter create new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).Name(route.Name).Path(route.Pattern).Handler(route.HandlerFunc)
	}
	return router
}
