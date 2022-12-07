package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Employee struct {
	Email string
	Phone string
	Role  string
}

type Task struct {
	Name string
}

type EmployeeTask struct {
	EmployeeID int
	TaskID     int
}

// Global Vars
var db *gorm.DB
var err error

func main() {
	// Load Environment Variables
	dialect := os.Getenv("DBDIALECT")
	host := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("DBUSER")
	dbName := os.Getenv("DBNAME")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable port=%s", host, user, dbName, dbPort)

	// Open Database Connection
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connection Successful.")
	}

	// Close Database Connection
	defer db.Close()

	// Since I'm not using the gorm.Model, skipping Migrations for now
	//db.AutoMigrate(&Employee{})
	//db.AutoMigrate(&Task{})
	//db.AutoMigrate(&EmployeeTask{})

}
