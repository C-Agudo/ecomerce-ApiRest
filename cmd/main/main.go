package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/C-Agudo/ecomerce-ApiRest/internal/database"
	"github.com/C-Agudo/ecomerce-ApiRest/internal/logs"

)

func main(){
	_ = logs.InitLogger()

	client := database.NewSqlClient(source: "root:root@/ecommerce")
}