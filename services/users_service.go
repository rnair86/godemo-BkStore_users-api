package services

import (
	"github.com/rnair86/godemo-BkStore_users-api/models/users"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userid int64) (*users.User,*errors.RestErr) {
	// if(userid <= 0) {
	// 	return nil, errors.NewBadRequestError("Invalid User ID")
	// }

	result := &users.User{Id: userid}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
