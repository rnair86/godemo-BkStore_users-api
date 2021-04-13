package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	UsersDb *sql.DB
)

func init() {
	const (
		username = "root"
		password = "test1234"
		protocol = "tcp"
		address  = "127.0.0.1"
		dbname   = "users_db"
	)

	datasourceName := fmt.Sprintf("%s:%s@%s(%s)/%s?", username, password, protocol, address, dbname) //username:password@protocol(address)/dbname?param=value
	var err error
	UsersDb, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}

	if err = UsersDb.Ping(); err != nil {
		panic(err)
	}

	//UsersDb = usersDB
	log.Println("database sussesfully configured")
}
