package dbadapter

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//Adapter Model
type Adapter struct {
	Table      *gorm.DB
	Connection *sql.DB
}

//Function to create connection to database
func (adp Adapter) New() Adapter {
	db_url := os.Getenv("DATABASE_URL")
	db_name := os.Getenv("DATABASE_NAME")
	if len(db_url) == 0 {
		db_url = "postgres://postgres:welcome1@localhost:5432/"
		db_name = "pokemon"
	}
	db, err := gorm.Open(postgres.Open(db_url+db_name), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	connection, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	return Adapter{Table: db, Connection: connection}
}
