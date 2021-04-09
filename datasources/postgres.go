package datasources

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // here
)

var Db *sql.DB

func Init() {
	tmpDB, err := sql.Open("postgres", "dbname=eventosdb user=postgres password=root host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	Db = tmpDB

	log.Println("Base de datos conectada...")
}
