package app

import (
	"github.com/rnair86/godemo-BkStore_users-api/controllers/ping"
	"github.com/rnair86/godemo-BkStore_users-api/controllers/users"
)

func mapUrls() {
	//Ping route
	router.GET("/ping", ping.Ping)

	//User Routes
	router.GET("/users/:user_id", users.Get)
	router.GET("/users/search", users.Search)

	router.POST("/users", users.Create)

	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)

	router.DELETE("/users/:user_id", users.Delete)

	router.GET("/internal/users/search", users.Search)

}
