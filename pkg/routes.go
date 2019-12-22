package calc

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		HomeLink,
	},
	Route{
		"sumHandler",
		"POST",
		"/v1/sum",
		SumHandler,
	},
	Route{
		"sqrtHandler",
		"POST",
		"/v1/sqrt",
		SqrtHandler,
	},
	Route{
		"factorialHandler",
		"POST",
		"/v1/factorial",
		FactorialHandler,
	},
	Route{
		"logHandler",
		"POST",
		"/v1/log",
		LogHandler,
	},
	Route{
		"isPrimeHandler",
		"POST",
		"/v1/isPrime",
		IsPrimeHandler,
	},
}
