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
	Route{
		"TodoIndexPost",
		"POST",
		"/todos",
		controllers.TodoIndexPost,
	},
}

func GetApi() Routes {
	return api
}
