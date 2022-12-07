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
	Name string `gorm:"column:task_name"`
}

type EmployeeTask struct {
	ID         int
	EmployeeID int
	TaskID     int
}

// EmployeeEmailTaskName serves the endpoint: getEmployeesByTaskName
type EmployeeEmailTaskName struct {
	EmployeeEmail string
	TaskName      string
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

	// API routes
	router := mux.NewRouter()

	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employee/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/employees/searchByTaskName/{searchterm}", getEmployeesByTaskName)
	router.HandleFunc("/employees/searchByPhone/{searchterm}", getEmployeesByPhoneNumber)

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")

	fmt.Println("Now Serving on localhost:8080")
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
	db.First(&employee, params["id"])
	json.NewEncoder(w).Encode(&employee)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	db.First(&task, params["id"])
	json.NewEncoder(w).Encode(&task)
}

// getEmployeesByTaskName returns a standard API response as JSON
// with a set of EmployeeEmail and TaskName
// where input 'searchterm' partially matches Task.Name, case-insensitive
// by using Inner Joins on the EmployeeTask table
func getEmployeesByTaskName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	/*
		// Leaving in these attempts to show thought development
		// First I wrote raw SQL that WORKS
		SELECT Employees.email AS EmployeeEmail, Tasks.Name AS TaskName
		FROM ((EmployeeTask
		INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID)
		INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID)
		WHERE LOWER(Tasks.Name) LIKE LOWER('%new%');

		// Next, I tried but failed several idiomatic Gorm Joins: https://gorm.io/docs/query.html#Joins
		// the following are succinct highlights
		db.Model(&EmployeeTask{}).Select(
		"Employees.email, Tasks.name").Joins(
		"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
		"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
		"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Scan(&employees)

		db.Select("Employees.email, Tasks.name").Joins(
		"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
		"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
		"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Find(&results)

		db.Table("employeetask").Select("Employees.email, Tasks.name").Joins(
		"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
		"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
		"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Scan(&results)

		// WORKS to return Employee Fields but I wanted Task Name and Gorm didn't like that
		db.Table("employeetask").Select("Employees.ID, Employees.email, Employees.Phone, Employees.Role").Joins(
		"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
		"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
		"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Scan(&employees)

		Finally decided to execute raw and extra raw from row data, then remarshal to json as shown below
	*/

	var employeeEmail string
	var taskName string
	var employeeEmailTaskNames []EmployeeEmailTaskName

	rows, err := db.Raw("SELECT Employees.email AS EmployeeEmail, Tasks.Name AS TaskName FROM ((EmployeeTask INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID) INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID) WHERE LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&employeeEmail, &taskName)
		// The following confirmed after several attempts I had the right data, commented for posterity
		//fmt.Println("EmployeeEmail:" + employeeEmail + " TaskName: " + taskName)
		result := EmployeeEmailTaskName{EmployeeEmail: employeeEmail, TaskName: taskName}
		employeeEmailTaskNames = append(employeeEmailTaskNames, result)
	}

	json.NewEncoder(w).Encode(&employeeEmailTaskNames)
}

// getEmployeesByPhoneNumber returns a standard API response as JSON
// with a set of Employee structs, all possible records
// where input 'searchterm' exactly matches Employee.Phone
func getEmployeesByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var employees []Employee
	db.Where("phone = ?", params["searchterm"]).Find(&employees)
	json.NewEncoder(w).Encode(&employees)
}
