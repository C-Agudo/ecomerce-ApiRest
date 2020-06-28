package database

import (
	"databade/sql"
	"database/sql/driver"
	"fmt"
)

type MySqlClient struct {

}

func NewSqlClient(source string) *sql.DB {
	db, err := sql.Open( driverName: "mysql", source)

	if err != nil {
		_= fmt.Error(format : "cannot create db tentat: %s", err.Error() )
		panic(v: "cannot create db")
	}

	return db
}