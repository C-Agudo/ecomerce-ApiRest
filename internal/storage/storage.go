package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// func ConnectToDB() *sql.DB {
// 	server := os.Getenv("GCP_HOSTNAME")
// 	if server == "" {
// 		server = "localhost"
// 	}
// 	connURL := fmt.Sprintf("root:root@tcp(localhost:3308/ecommerce?sslmode=disable", server)
// 	db, err := sql.Open("postgres", connURL)
// 	if err != nil {
// 		log.Fatalf("Failed to connect to DB via %s: %v", connURL, err)
// 	}
// 	if err = db.Ping(); err != nil {
// 		log.Fatalf("Failed to ping DB via %s: %v", connURL, err)
// 	}
// 	log.Println("Connected to DB")
// 	return db
// }
func ConnectToDB() (db *sql.DB) {

	usuario := "root"
	pass := "root"
	host := "tcp(127.0.0.1:3308)"
	nombreBaseDeDatos := "ecommerce"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil
	}
	return db
}
