package users

import (
	//"encoding/json"
	"fmt"
	"strconv"

	"github.com/rnair86/godemo-BkStore_users-api/services"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"

	//"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rnair86/godemo-BkStore_users-api/models/users"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	fmt.Println("error: ",err)
	// 	//TODO: Handle error
	// 	return
	// }
	// if err:= json.Unmarshal(bytes,&user); err!=nil {
	// 	fmt.Println("error: ",err)
	// 	//TODO: Handle JSON Conv error
	// 	return
	// }
	//Replace with
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		fmt.Println("error: ", restErr)
		c.JSON(restErr.Status, restErr)
		//TODO: Handle Json error
		return
	}

	fmt.Printf("%+v\n", user)
	//fmt.Println(user)

	result, saveerr := services.CreateUser(user)
	if saveerr != nil {
		fmt.Println("error: ", saveerr)
		c.JSON(saveerr.Status, saveerr)
		return
	}

	fmt.Printf("%+v\n", result)

	//c.String(http.StatusNotImplemented,"Not implemented yet Check back soon!!")
	c.JSON(http.StatusCreated, result)
}
func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("Invalid user_id")
		c.JSON(err.Status, err)
		return
	}

	user, geterr := services.GetUser(userId)
	if geterr != nil {
		fmt.Println("error: ", geterr)
		c.JSON(geterr.Status, geterr)
		return
	}
	c.JSON(http.StatusOK, user)

	//c.String(http.StatusNotImplemented, "Not implemented yet Check back soon!!")
}
func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Not implemented yet Check back soon!!")
}
