package services

import (
	"github.com/rnair86/godemo-BkStore_users-api/models/users"
	"github.com/rnair86/godemo-BkStore_users-api/utils/crypto_utils"
	"github.com/rnair86/godemo-BkStore_users-api/utils/date_utils"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMD5(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userid int64) (*users.User, *errors.RestErr) {

	result := &users.User{Id: userid}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}

	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}

	if upderr := current.Update(); err != nil {
		return nil, upderr
	}
	return current, nil

}

func DeleteUser(userId int64) *errors.RestErr {

	user := &users.User{Id: userId}
	return user.Delete()
}

func FindByStatus(status string) (users.Users, *errors.RestErr) {
	userdao := &users.User{}
	return userdao.FindByStatus(status)
}
