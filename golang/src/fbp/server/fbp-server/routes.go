package fbpserver

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	//WEBSOCKETS
	Route{
		"webSocketHandler",
		"GET",
		"/",
		webSocketHandler,
	},
}
