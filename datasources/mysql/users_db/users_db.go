package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rnair86/godemo-BkStore_users-api/logger"
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
		logger.Error("error opening mysql db connection", err)
		panic(err)
	}

	if err = UsersDb.Ping(); err != nil {
		logger.Error("error pinging mysql database", err)
		panic(err)
	}

	//UsersDb = usersDB
	logger.Info("database successfully configured")
}
