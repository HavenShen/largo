package routes

import (
	"github.com/HavenShen/largo/controllers"
)

var api = Routes{
	Route{
		"Index",
		"GET",
		"/",
		controllers.Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		controllers.TodoIndex,
	},
}

func GetApi() Routes {
	return api
}
