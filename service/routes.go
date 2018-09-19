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
		"GetAge",
		"POST",
		"/getage/{age}",
		handler.GetAge,
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
		"Login",
		"POST",
		"/login",
		handler.Login,
	},
	Route{
		"AddArticle",
		"POST",
		"/addArticle",
		handler.ValidateTokenMiddleware(handler.AddArticle),
	},
	Route{
		"UpdateArticle",
		"POST",
		"/updateArticle",
		handler.ValidateTokenMiddleware(handler.UpdateArticle),
	},
	Route{
		"DeleteArticle",
		"POST",
		"/deleteArticle",
		handler.ValidateTokenMiddleware(handler.DeleteArticle),
	},
	Route{
		"FindArticle",
		"POST",
		"/findArticle",
		handler.ValidateTokenMiddleware(handler.FindArticle),
	},
}
