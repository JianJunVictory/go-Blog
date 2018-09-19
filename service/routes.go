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

// Routes defined
type Routes []Route

var routes = Routes{
	Route{
		"GetName",
		"POST",
		"/getname",
		handler.ValidateTokenMiddleware(handler.GetName),
	},
	Route{
		"GetAge",
		"POST",
		"/getage/{age}",
		handler.GetAge,
	},
	Route{
		"Login",
		"POST",
		"/login",
		handler.Login,
	},
	Route{
		"Register",
		"POST",
		"/register",
		handler.Register,
	},
	Route{
		"ActiveAccount",
		"POST",
		"/activeAccount",
		handler.ActiveAccount,
	},
	Route{
		"AddArticle",
		"POST",
		"/addArticle",
		handler.ValidateTokenMiddleware(handler.AddArticle),
	},
}
