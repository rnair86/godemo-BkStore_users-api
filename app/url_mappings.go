package app

import (
	"github.com/rnair86/godemo-BkStore_users-api/controllers/ping"
	"github.com/rnair86/godemo-BkStore_users-api/controllers/users"
)

func mapUrls() {
	//Ping route
	router.GET("/ping", ping.Ping)

	//User Routes
	router.GET("/users/:user_id",users.GetUser)
	router.GET("/users/search",users.SearchUser)
	router.POST("/users",users.CreateUser)

}