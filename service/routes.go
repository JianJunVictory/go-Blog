package service

import (
	"net/http"

	"github.com/go-Blog/handler"
)

// Route route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes define
type Routes []Route

var routes = Routes{
	Route{
		"GetName",
		"POST",
		"/getname",
		handler.GetName,
	},
	Route{
		"GetAge",
		"POST",
		"/getage/{age}",
		handler.GetAge,
	},
}
