package users

import (
	"github.com/rnair86/godemo-BkStore_users-api/logger"
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
		logger.Error("error while mapping received user Json", err)
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveerr := services.UsersService.CreateUser(user)
	if saveerr != nil {
		c.JSON(saveerr.Status, saveerr)
		return
	}

	//logger.Info(fmt.Sprintf("%+v\n", result))

	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func Get(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status, idErr)
		return
	}

	user, geterr := services.UsersService.GetUser(userId)
	if geterr != nil {
		//fmt.Println("error: ", geterr)
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
		logger.Error("error while mapping user Json", err)
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, upderr := services.UsersService.UpdateUser(isPartial, user)

	if upderr != nil {
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

	if err := services.UsersService.DeleteUser(userId); err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {

	status := c.Query("status")
	result, searcherr := services.UsersService.SearchUser(status)
	if searcherr != nil {
		c.JSON(searcherr.Status, searcherr)
		return
	}

	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))

}

func getUserId(userIdParam string) (int64, *errors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		logger.Error("error while parsing userid", userErr)
		return 0, errors.NewBadRequestError("Invalid user_id")
	}
	return userId, nil
}
