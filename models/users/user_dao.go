package users

import (
	"fmt"

	"github.com/rnair86/godemo-BkStore_users-api/datasources/mysql/users_db"
	"github.com/rnair86/godemo-BkStore_users-api/utils/errors"
	"github.com/rnair86/godemo-BkStore_users-api/utils/mysql_utils"
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
		return errors.NewInternalServerError(dberr.Error())
	}
	defer stmt.Close() // **Important**

	fmt.Printf("%+v\n", user)

	insrtRslt, insrtErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated,user.Password,user.Status)

	if insrtErr != nil {
		return mysql_utils.ParseError(insrtErr)
	}

	userid, irErr := insrtRslt.LastInsertId()
	if irErr != nil {
		return mysql_utils.ParseError(irErr)
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

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		return mysql_utils.ParseError(err)
	}
	return nil

}

func (user *User) Update() *errors.RestErr {
	stmt, dberr := users_db.UsersDb.Prepare(queryUpdateUser)
	if dberr != nil {
		return errors.NewInternalServerError(dberr.Error())
	}
	defer stmt.Close()

	_, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Id)
	if err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) Delete() *errors.RestErr {

	stmt, dberr := users_db.UsersDb.Prepare(queryDeleteUser)
	if dberr != nil {
		return errors.NewInternalServerError(dberr.Error())
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Id); err != nil {
		return mysql_utils.ParseError(err)
	}

	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, dberr := users_db.UsersDb.Prepare(queryFindUserByStatus)

	if dberr != nil {
		return nil, errors.NewInternalServerError(dberr.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)

		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundRequestError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil

}
