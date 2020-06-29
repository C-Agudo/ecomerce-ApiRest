package database

import (
	"database/sql"
	"fmt"

	"github.com/C-Agudo/ecomerce-ApiRest/internal/logs"
	_ "github.com/go-sql-driver/mysql"
)

type MySqlClient struct {
	*sql.DB
}

// func NewSqlClient(source string) *MySqlClient {
// 	db, err := sql.Open("mysql", source)

// 	if err != nil {
// 		logs.Log().Errorf("cannot create db")
// 	}

// 	err = db.Ping()

// 	if err != nil {
// 		logs.Log().Warn("cannot conect to mysql")
// 	}
// 	return &MySqlClient{db}
// }
// type MySqlClient struct {
// 	*sql.DB
// }

func NewSqlClient() *MySqlClient {
	usuario := "root"
	pass := "root"
	host := "tcp(127.0.0.1:3308)"
	nombreBaseDeDatos := "ecommerce"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		logs.Log().Errorf("cannot create db tentat: %s", err.Error())
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		logs.Log().Warn("cannot connect to mysql!")
	}

	return &MySqlClient{db}

	// db, err := sql.Open("mysql", source)

	// if err != nil {
	// 	logs.Log().Errorf("cannot create db tentat: %s", err.Error())
	// 	panic(err)
	// }

	// err = db.Ping()

	// if err != nil {
	// 	logs.Log().Warn("cannot connect to mysql!")
	// }

	// return &MySqlClient{db}
}

func (c *MySqlClient) ViewStats() sql.DBStats {
	return c.Stats()
}
