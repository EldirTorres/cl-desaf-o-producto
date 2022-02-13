package config

import (
	"cl.falabella.product/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{

	Route{
		"HealtCheck",
		"GET",
		"/",
		controller.HealtCheck,
	},
	Route{
		"Saveproduct",
		"POST",
		"/product",
		controller.Saveproduct,
	},
}