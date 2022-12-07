package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Employee struct {
	ID    int
	Email string
	Phone string
	Role  string
}

type Task struct {
	ID   int
	Name string
}

type EmployeeTask struct {
	ID         int
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

	// Since I'm not using the gorm.Model, skip Migrations
	//db.AutoMigrate(&Employee{})
	//db.AutoMigrate(&Task{})
	//db.AutoMigrate(&EmployeeTask{})

	// API routes
	router := mux.NewRouter()

	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/tasks", getTasks).Methods("GET")

	http.ListenAndServe(":8080", router)

}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	var employees []Employee
	db.Find(&employees)
	json.NewEncoder(w).Encode(&employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var employee Employee

	// Find the first record matching the condition, ordered by Primary Key
	db.First(&employee, params["id"])

	json.NewEncoder(w).Encode(&employee)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}
