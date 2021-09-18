package users

import (
	//"encoding/json"
	"fmt"
	"github.com/rnair86/godemo-BkStore_users-api/services"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
	"strconv"

	//"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rnair86/godemo-BkStore_users-api/models/users"
)

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		fmt.Println("error: ", restErr)
		c.JSON(restErr.Status, restErr)
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

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}
func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, geterr := services.GetUser(userId)
	if geterr != nil {
		fmt.Println("error: ", geterr)
		c.JSON(geterr.Status, geterr)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	//get userid from request
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		fmt.Println("error: ", restErr)
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	fmt.Printf("%+v\n", user)

	isPartial := c.Request.Method == http.MethodPatch

	result, upderr := services.UpdateUser(isPartial, user)

	if upderr != nil {
		fmt.Println("error: ", upderr)
		c.JSON(upderr.Status, upderr)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	if err := services.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {

	status := c.Query("status")
	result, searcherr := services.FindByStatus(status)
	if searcherr != nil {
		c.JSON(searcherr.Status, searcherr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {

		return 0, errors.NewBadRequestError("Invalid user_id")
	}
	return userId, nil
}
