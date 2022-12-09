package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux" // Serves API Routes

	"gorm.io/driver/postgres" // Enables gorm to use postgres driver
	"gorm.io/gorm"            // Facilitates postgres db queries
)

// Employee matches the Employee Table
type Employee struct {
	ID    int
	Email string
	Phone string
	Role  string
}

// Task matches the Task Table
type Task struct {
	ID   int
	Name string
}

// EmployeeTask matches the EmployeeTask Table
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

// main opens the database connection, starts an http lisetener, and
// keeps running until the process is terminated by the user/system
func main() {
	// Load Environment Variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)

	// Connection Loop, on failure try try again
	// Works for this small example; for prod, I'd use a better pattern, ie. Connection Pool
	for {

		// Open Database Connection
		db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("DB Connection Successful.")
			break
		}

		fmt.Print(".") // small progress indicator
		time.Sleep(1 * time.Second)
		continue
	}

	// API routes
	router := mux.NewRouter()

	router.HandleFunc("/employees", getEmployees).Methods("GET").Headers("TrailHead-token", "pa$$word")
	router.HandleFunc("/employee/{id}", getEmployee).Methods("GET").Headers("TrailHead-token", "pa$$word")
	router.HandleFunc("/employees/searchByTaskName/{searchterm}", getEmployeesByTaskName).Headers("TrailHead-token", "pa$$word")
	router.HandleFunc("/employees/searchByPhone/{searchterm}", getEmployeesByPhoneNumber).Headers("TrailHead-token", "pa$$word")

	router.HandleFunc("/tasks", getTasks).Methods("GET").Headers("TrailHead-token", "pa$$word")
	router.HandleFunc("/task/{id}", getTask).Methods("GET").Headers("TrailHead-token", "pa$$word")

	fmt.Println("Now Serving on localhost:8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}

func hasCorrectHeaders(r *http.Request) bool {
	return hasTrailHeadToken(r)
}

func hasTrailHeadToken(r *http.Request) bool {
	val := r.Header.Get("TrailHead-token")
	return val == "pa$$word"
}

// getEmployees returns a standard API response as JSON
// with a set of Employee structs, all records from the Employee table
// no input
func getEmployees(w http.ResponseWriter, r *http.Request) {
	if !hasCorrectHeaders(r) {
		fmt.Fprintf(w, "Permission Denied.\n")
		log.Println("Request failed to have correct headers.")
		return
	}

	var employees []Employee
	db.Find(&employees)
	json.NewEncoder(w).Encode(&employees)
}

// getEmployee returns a standard API response as JSON
// with one Employee struct, a single record
// where input 'id' exactly matches the auto-generated int for the record
func getEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var employee Employee
	db.First(&employee, params["id"])
	json.NewEncoder(w).Encode(&employee)
}

// getTasks returns a standard API response as JSON
// with a set of Task structs, all records from the Tasks table
// no input
func getTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	db.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

// getTask returns a standard API response as JSON
// with one Task struct, a single record
// where input 'id' exactly matches the auto-generated int for the record
func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	db.First(&task, params["id"])
	json.NewEncoder(w).Encode(&task)
}

// getEmployeesByTaskName returns a standard API response as JSON
// with a set employeeEmailTaskNames structs, multiple records
// where input 'searchterm' partially matches Task.Name, case-insensitive
// by using Inner Joins on the EmployeeTask table
func getEmployeesByTaskName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

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
