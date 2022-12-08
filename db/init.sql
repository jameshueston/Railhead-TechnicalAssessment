-- CREATE TABLES

CREATE TABLE Employees (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	email VARCHAR(50) NOT NULL,
	phone VARCHAR(50) NOT NULL,
	role VARCHAR(50) NOT NULL
);

CREATE TABLE Tasks (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	name VARCHAR(100) NOT NULL
);

CREATE TABLE employeetask (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	employeeid INT NOT NULL REFERENCES employees(id),
	taskid INT NOT NULL REFERENCES tasks(id)
);


-- ADD CONSTRAINTS

ALTER TABLE employees
ADD CONSTRAINT role_constraint
CHECK (role = 'supervisor' OR role = 'worker');


-- SEED TABLE: Employees

INSERT INTO Employees (email, phone, role)
VALUES ('ebanks@mlb.org', '6303236630', 'supervisor');

INSERT INTO Employees (email, phone, role)
VALUES ('retiredryno@windycityallstars.org','1009071984','supervisor');

INSERT INTO Employees (email, phone, role)
VALUES ('sweetswingbilly@hofoutfielders.com','2963921353','supervisor');

INSERT INTO Employees (email, phone, role)
VALUES ('bestpitcherjenkins@wrigleyfield.com','1671321971', 'supervisor');

INSERT INTO Employees (email, phone, role)
VALUES ('sosa@cubs.mlb','5454966148','supervisor');

INSERT INTO Employees (email, phone, role)
VALUES ('jhueston@railheadcorp.com', '7088445500', 'worker');

INSERT INTO Employees (email, phone, role)
VALUES ('benchtester@railheadcorp.com', '7088445500', 'worker');

INSERT INTO Employees (email, phone, role)
VALUES ('support@railheadcorp.com', '7088445500', 'worker');

INSERT INTO Employees (email, phone, role)
VALUES ('quality@railheadcorp.com', '7088445500', 'worker');

INSERT INTO Employees (email, phone, role)
VALUES ('systems@railheadcorp.com', '7088445500', 'worker');


-- SEED TABLE: Tasks

INSERT INTO Tasks (name) VALUES ('Find the Right People');
INSERT INTO Tasks (name) VALUES ('Tailor Jobs to Fit New Hires');
INSERT INTO Tasks (name) VALUES ('Make Fleetwide Trackers Better');
INSERT INTO Tasks (name) VALUES ('Create Back Office and Web Apps Customers Need');
INSERT INTO Tasks (name) VALUES ('Tailor New Product to Small Customers Efficiently');


-- SEED TABLE: EmployeeTask

-- Assign Tasks to Supervisors
INSERT INTO employeetask (employeeid, taskid)
VALUES (1,1);

INSERT INTO employeetask (employeeid, taskid)
VALUES (1,2);

-- There was no requirement for every employee to have a task, so prove that it's possible.
-- Since Ryno is retired (EmployeeID = 2), don't assign him a task; Intentionally commented:
-- INSERT INTO employeetask (employeeid, taskid)
-- VALUES (2,2);

INSERT INTO employeetask (employeeid, taskid)
VALUES (3,1);

INSERT INTO employeetask (employeeid, taskid)
VALUES (3,3);

INSERT INTO employeetask (employeeid, taskid)
VALUES (4,4);

INSERT INTO employeetask (employeeid, taskid)
VALUES (5,5);

-- Assign Tasks to Workers
INSERT INTO employeetask (employeeid, taskid)
VALUES (6,4);

INSERT INTO employeetask (employeeid, taskid)
VALUES (6,5);

INSERT INTO employeetask (employeeid, taskid)
VALUES (7,3);

INSERT INTO employeetask (employeeid, taskid)
VALUES (7,5);

INSERT INTO employeetask (employeeid, taskid)
VALUES (8,3);

INSERT INTO employeetask (employeeid, taskid)
VALUES (8,4);

INSERT INTO employeetask (employeeid, taskid)
VALUES (9,3);

INSERT INTO employeetask (employeeid, taskid)
VALUES (10,3);

INSERT INTO employeetask (employeeid, taskid)
VALUES (10,4);

-- intentional end