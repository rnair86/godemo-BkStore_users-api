package users

import (
	"github.com/rnair86/godemo-BkStore_users-api/logger"
	"strings"

	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"-"`
}

type Users []User

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("Invalid Password")
	}
	logger.Info("User Validated!!")
	return nil

}
