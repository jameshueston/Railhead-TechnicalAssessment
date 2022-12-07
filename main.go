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

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by EmployeeTask to `employeetask`
func (EmployeeTask) TableName() string {
	return "employeetask"
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
	router.HandleFunc("/employees/searchByTaskName/{searchterm}", getEmployeesByTaskName)

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")

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

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var task Task
	db.First(&task, params["id"])
	json.NewEncoder(w).Encode(&task)
}

type EmployeeEmailTaskName struct {
	EmployeeEmail string
	TaskName      string
}

func getEmployeesByTaskName(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)

	//var employees []Employee
	//var task Task
	//var employeeTasks []EmployeeTask
	//var employeeEmailTaskNames []EmployeeEmailTaskName

	// "searchterm" accepts partial name matches, case-insensitive
	// db.First(&task, params["searchterm"])
	// now the task we want should be in &task with task.ID searchable in EmployeeTask

	//whereConditionFormatted := fmt.Sprintf("LOWER(Tasks.Name) LIKE LOWER('\%%s\%')", params["searchterm"])
	db.LogMode(true)

	var results []map[string]interface{}
	//db.Table("users").Find(&results)
	var EmployeeEmail string
	var TaskName string

	//db.Raw("SELECT Employees.email AS EmployeeEmail, Tasks.Name AS TaskName FROM ((EmployeeTask INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID) INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID) WHERE LOWER(Tasks.Name) LIKE LOWER('%new%')").Scan(&results)
	rows, err := db.Raw("SELECT Employees.email AS EmployeeEmail, Tasks.Name AS TaskName FROM ((EmployeeTask INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID) INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID) WHERE LOWER(Tasks.Name) LIKE LOWER('%new%')").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&EmployeeEmail, &TaskName)
		fmt.Println("EmployeeEmail:" + EmployeeEmail + " TaskName: " + TaskName)
	}

	/*
		db.Table("employeetask").Select("Employees.email, Tasks.name").Joins(
			"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
			"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
			"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Scan(&results)

		/*
			db.Select("Employees.email, Tasks.name").Joins(
			"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
			"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
			"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Find(&results)

			db.Model(&EmployeeTask{}).Select(
			"Employees.email, Tasks.name").Joins(
			"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
			"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
			"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Scan(&employees)

			// WORKS - Returning Employee Fields but not yet Task Name field
			db.Table("employeetask").Select("Employees.ID, Employees.email, Employees.Phone, Employees.Role").Joins(
			"INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID").Joins(
			"INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID").Where(
			"LOWER(Tasks.Name) LIKE LOWER('%" + params["searchterm"] + "%')").Scan(&employees)

			// WORKS - Search Task name by keyword, ignoring case
			SELECT Employees.email AS EmployeeEmail, Tasks.Name AS TaskName
			FROM ((EmployeeTask
			INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID)
			INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID)
			WHERE LOWER(Tasks.Name) LIKE LOWER('%new%');




				db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
				// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

				rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
				for rows.Next() {
					...
				}

				db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

				// multiple joins with parameter
				db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)

	*/
	fmt.Println(&results)
	json.NewEncoder(w).Encode(&results)
}
