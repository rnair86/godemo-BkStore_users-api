package users

import (
	"strings"

	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	LastName    string `json:"last_name"`
	FirstName   string `json:"first_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid Email address")
	}
	return nil

}
