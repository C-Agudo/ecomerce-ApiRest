package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open("mysql", source)

	if err != nil {
		_ = fmt.Errorf("cannot create db tentat: %s", err.Error())
		panic(err)
	}

	if err != nil {
		logs.Log().Warn(args...:"cannot conect to mysql")
	} 
	return &MySqlClient(DB:db)
}
