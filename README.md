# Getting Started

1. Clone the Project
    - `INSERT CLONE HERE`

1. Start Services
    - `docker compose up`
    - For expected output, see [docs/expected_output/docker_compose_up](docs/expected_output/docker_compose_up).
    
1. Execute Queries (see below)

1. When finished, clean up:
    - `docker compose rm --force`
    - `docker volume rm technicalassessment_db`

This was tested on `docker v20.10.21`.

# Example Queries
    
The service outputs JSON.

Examples below use the small command-line utility `jq` to pretty print -- highly recommend! You may [Install jq](https://stedolan.github.io/jq/download/) or  remove `| jq` from queries before executing.

`curl --silent` removes download status and is provided on examples for clean output when piping to `jq`.

Per requirements, each successful query needs the Header `TrailHead-token: pa$$word`. Examples below use `curl` and escape the dollar signs with backslashes.

Two queries were required:

**Query 5: Search by a phone number to look up employees**

**Query 6: Search by a task name to get all employees working on it.**

In order to discover searchable terms for those queries, additional queries are provided below showing data available.

## Query 1: Get All Employees

`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/employees | jq`                

```json
[
  {
    "ID": 1,
    "Email": "ebanks@mlb.org",
    "Phone": "6303236630",
    "Role": "supervisor"
  },
  {
    "ID": 2,
    "Email": "retiredryno@windycityallstars.org",
    "Phone": "1009071984",
    "Role": "supervisor"
  },
  {
    "ID": 3,
    "Email": "sweetswingbilly@hofoutfielders.com",
    "Phone": "2963921353",
    "Role": "supervisor"
  },
  {
    "ID": 4,
    "Email": "bestpitcherjenkins@wrigleyfield.com",
    "Phone": "1671321971",
    "Role": "supervisor"
  },
  {
    "ID": 5,
    "Email": "sosa@cubs.mlb",
    "Phone": "5454966148",
    "Role": "supervisor"
  },
  {
    "ID": 6,
    "Email": "jhueston@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 7,
    "Email": "benchtester@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 8,
    "Email": "support@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 9,
    "Email": "quality@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 10,
    "Email": "systems@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  }
]
```

## Query 2 - Get Employee by ID
`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/employee/1 | jq`

```json
{
  "ID": 1,
  "Email": "ebanks@mlb.org",
  "Phone": "6303236630",
  "Role": "supervisor"
}
```

## Query 3: Get All Tasks

`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/tasks | jq`

```json
[
  {
    "ID": 1,
    "Name": "Find the Right People"
  },
  {
    "ID": 2,
    "Name": "Tailor Jobs to Fit New Hires"
  },
  {
    "ID": 3,
    "Name": "Make Fleetwide Trackers Better"
  },
  {
    "ID": 4,
    "Name": "Create Back Office and Web Apps Customers Need"
  },
  {
    "ID": 5,
    "Name": "Tailor New Product to Small Customers Efficiently"
  }
]
```

## Query 4: Get Task by ID
`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/task/1 | jq`

```json
{
  "ID": 1,
  "Name": "Find the Right People"
}
```


## Query 5: Search by a phone number to look up employees

### Example 1 - Single matching record

`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/employees/searchByPhone/5454966148 | jq`

```json
[
  {
    "ID": 5,
    "Email": "sosa@cubs.mlb",
    "Phone": "5454966148",
    "Role": "supervisor"
  }
]
```

### Example 2 - Multiple matching records
`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/employees/searchByPhone/7088445500 | jq`

```json
[
  {
    "ID": 6,
    "Email": "jhueston@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 7,
    "Email": "benchtester@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 8,
    "Email": "support@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 9,
    "Email": "quality@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  },
  {
    "ID": 10,
    "Email": "systems@railheadcorp.com",
    "Phone": "7088445500",
    "Role": "worker"
  }
]
```

## Query 6: Search by a task name to get all employees working on it.

A case-insensitive search on Task.Name for {searchterm}

`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/employees/searchByTaskName/{searchterm} | jq`

### Example 1 - where 'to' is `{searchterm}`

`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/employees/searchByTaskName/to | jq`

```json
[
  {
    "EmployeeEmail": "ebanks@mlb.org",
    "TaskName": "Tailor Jobs to Fit New Hires"
  },
  {
    "EmployeeEmail": "bestpitcherjenkins@wrigleyfield.com",
    "TaskName": "Create Back Office and Web Apps Customers Need"
  },
  {
    "EmployeeEmail": "sosa@cubs.mlb",
    "TaskName": "Tailor New Product to Small Customers Efficiently"
  },
  {
    "EmployeeEmail": "jhueston@railheadcorp.com",
    "TaskName": "Create Back Office and Web Apps Customers Need"
  },
  {
    "EmployeeEmail": "jhueston@railheadcorp.com",
    "TaskName": "Tailor New Product to Small Customers Efficiently"
  },
  {
    "EmployeeEmail": "benchtester@railheadcorp.com",
    "TaskName": "Tailor New Product to Small Customers Efficiently"
  },
  {
    "EmployeeEmail": "support@railheadcorp.com",
    "TaskName": "Create Back Office and Web Apps Customers Need"
  },
  {
    "EmployeeEmail": "systems@railheadcorp.com",
    "TaskName": "Create Back Office and Web Apps Customers Need"
  }
]
```

### Example 2 - where 'the' is `{searchterm}`

`curl --silent --header "TrailHead-token:pa\$\$word" localhost:8080/employees/searchByTaskName/the | jq`

```json
[
  {
    "EmployeeEmail": "ebanks@mlb.org",
    "TaskName": "Find the Right People"
  },
  {
    "EmployeeEmail": "sweetswingbilly@hofoutfielders.com",
    "TaskName": "Find the Right People"
  }
]
```
