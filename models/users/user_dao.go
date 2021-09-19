package users

import (
	"fmt"
	"github.com/rnair86/godemo-BkStore_users-api/logger"

	"github.com/rnair86/godemo-BkStore_users-api/datasources/mysql/users_db"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
)

const (
	queryInsertUser       = "INSERT INTO users(first_name,last_name,email,date_created,password,status) VALUES(?,?,?,?,?,?);"
	queryGetUserbyId      = "SELECT id,first_name,last_name,email,date_created, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id,first_name,last_name,email,date_created,status FROM users WHERE status=?"
)

func (user *User) Save() *errors.RestErr {
	stmt, dberr := users_db.UsersDb.Prepare(queryInsertUser)
	if dberr != nil {
		logger.Error("error when trying to prepare Save user statement", dberr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close() // **Important**
	insrtRslt, insrtErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)

	if insrtErr != nil {
		logger.Error("error when trying to Save user", insrtErr)
		return errors.NewInternalServerError("database error")
		//return mysql_utils.ParseError(insrtErr)
	}

	userid, irErr := insrtRslt.LastInsertId()
	if irErr != nil {
		logger.Error("error when trying to get last inserted userid", irErr)
		return errors.NewInternalServerError("database error")
		//return mysql_utils.ParseError(irErr)
	}

	user.Id = userid

	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, dberr := users_db.UsersDb.Prepare(queryGetUserbyId)
	if dberr != nil {
		logger.Error("error when trying to prepare get user statement", dberr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to get user", err)
		return errors.NewInternalServerError("database error")
		//return mysql_utils.ParseError(err)
	}
	return nil

}

func (user *User) Update() *errors.RestErr {
	stmt, dberr := users_db.UsersDb.Prepare(queryUpdateUser)
	if dberr != nil {
		logger.Error("error when trying to prepare update user statement", dberr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	_, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
	if err != nil {
		logger.Error("error when trying to Update user", err)
		return errors.NewInternalServerError("database error")
		//return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {

	stmt, dberr := users_db.UsersDb.Prepare(queryDeleteUser)
	if dberr != nil {
		logger.Error("error when trying to prepare delete user statement", dberr)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to Delete user", err)
		return errors.NewInternalServerError("database error")
		//return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, dberr := users_db.UsersDb.Prepare(queryFindUserByStatus)

	if dberr != nil {
		logger.Error("error when trying to prepare Search user statement", dberr)
		return nil, errors.NewInternalServerError("database error")

	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to execute Search user statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying read Search user", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		logger.Warn(fmt.Sprintf("no users matching status %s", status))
		return nil, errors.NewNotFoundRequestError("Cannot find matching users")
	}

	return results, nil

}
