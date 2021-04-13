package users

import (
	"fmt"
	"strings"

	//"log"

	"github.com/go-sql-driver/mysql"
	"github.com/rnair86/godemo-BkStore_users-api/datasources/mysql/users_db"
	"github.com/rnair86/godemo-BkStore_users-api/utils/date_utils"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
)

const (
	queryInsertUser  = "INSERT INTO users(first_name,last_name,email,date_created) VALUES(?,?,?,?);"
	indexUniqueEmail = "email_UNIQUE"
	queryGetUserbyId = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
	noRowsinResult   = "no rows in result set"
)

func (user *User) Save() *errors.RestErr {
	stmt, dberr := users_db.UsersDb.Prepare(queryInsertUser)
	if dberr != nil {
		return errors.NewInternalServerError(dberr.Error())
	}
	defer stmt.Close() // **Important**

	user.DateCreated = date_utils.GetNowString()

	insrtRslt, insrtErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if insrtErr != nil {
		sqlErr,ok := insrtErr.(*mysql.MySQLError)
		if !ok {
			return errors.NewInternalServerError(fmt.Sprintf("Error when saving User %s",insrtErr.Error()))
		}
		switch sqlErr.Number{
		case 1062:
			return errors.NewBadRequestError(fmt.Sprintf("Email address %s already in exists", user.Email))
		}
		
		//fmt.Println(sqlErr)
		// if strings.Contains(insrtErr.Error(), indexUniqueEmail) {
		// 	return errors.NewBadRequestError(fmt.Sprintf("Email address %s already in exists", user.Email))
		// }
		
	}

	userid, irErr := insrtRslt.LastInsertId()
	if irErr != nil {
		return errors.NewInternalServerError(irErr.Error())
	}

	user.Id = userid

	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, dberr := users_db.UsersDb.Prepare(queryGetUserbyId)
	if dberr != nil {
		return errors.NewInternalServerError(dberr.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), noRowsinResult) {
			return errors.NewNotFoundRequestError(fmt.Sprintf("No users exists for Id %d", user.Id))
		}

		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to get  user of id %d : %s", user.Id, err.Error()))
	}
	return nil

}
