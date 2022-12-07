SELECT Orders.OrderID, Customers.CustomerName, Shippers.ShipperName
FROM ((Orders
INNER JOIN Customers ON Orders.CustomerID = Customers.CustomerID)
INNER JOIN Shippers ON Orders.ShipperID = Shippers.ShipperID);


-- WORKS, showing all employee emails that have EmployeeID and TaskID in the EmployeeTask table
SELECT Employees.email
FROM ((EmployeeTask
INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID)
INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID);

-- WORKS
SELECT Employees.email, Tasks.name
FROM ((EmployeeTask
INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID)
INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID);

-- WORKS
SELECT Employees.email, Tasks.name
FROM ((EmployeeTask
INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID)
INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID)
WHERE Tasks.ID = 1;

-- WORKS
SELECT Employees.email, Tasks.name
FROM ((EmployeeTask
INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID)
INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID)
WHERE Tasks.Name = 'Find the Right People';

-- WORKS -- exactly what he asked for
SELECT Employees.email
FROM ((EmployeeTask
INNER JOIN Employees ON EmployeeTask.EmployeeID = Employees.ID)
INNER JOIN Tasks ON EmployeeTask.TaskID = Tasks.ID)
WHERE Tasks.Name = 'Find the Right People';

-- Search Task name by keyword, ignoring case
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
