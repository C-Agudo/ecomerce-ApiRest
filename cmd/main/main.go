package main

import (
	"github.com/C-Agudo/ecomerce-ApiRest/internal/database"
	"github.com/C-Agudo/ecomerce-ApiRest/internal/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationsRootFolder     = "file://migrations"
	migrationsScriptsVersion = 3
)

func main() {
	_ = logs.InitLogger()

	db := database.NewSqlClient()
	doMigrate(db, "ecommerce")
	defer db.Close()

	mux := Routes()
	server := NewServer(mux)
	server.Run()

}

// func main() {
// 	_ = logs.InitLogger()

// 	client := database.NewSqlClient("root:root@tcp(localhost:3308/ecommerce")
// 	doMigrate(client, "ecommerce")
// }

func doMigrate(database *database.MySqlClient, dbName string) {
	driver, _ := migration.WithInstance(database.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationsRootFolder,
		dbName,
		driver,
	)
	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("current migrations version in &d", current)
	err = m.Migrate(migrationsScriptsVersion)
	if err != nil && err.Error() == "no change" {
		logs.Log().Info("no migration needed")
	}
}
