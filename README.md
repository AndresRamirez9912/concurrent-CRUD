# Vozy-tech-evaluation

This project is a task management system where you can Add, Update, Get, and Delete tasks using SQLite DB. It uses concurrency to increase the scalability of the project and handle a large number of requests simultaneously.

| Field       | Use                                                            |
| ----------- | -------------------------------------------------------------- |
| TItle       | Assign a name or title of the task                             |
| Description | Description of the task or instructions                        |
| State       | Current state of the task (pendiente, en progreso, completada) |

### Dependencies

| Dependency | Version  |
| ---------- | -------- |
| go         | v1.20    |
| Gin        | v1.9.1   |
| sqlx       | v1.3.5   |
| go-sqlite3 | v1.14.22 |

## Getting started

Before starting the project, make sure you have SQLite installed on your PC.

### Install DB

#### Linux devices

Update the package management:

```bash
$ sudo apt update
```

Install SQLite

```bash
$ sudo apt install sqlite3
```

Verify the installation by checking the version:

```bash
$ sqlite3 --version
```

### 1. Install Dependencies

In order to install the project's dependencies, run this command:

```bash
$ go get
```

### 2. Initialyze the DB in your local

The project contains a file called "init.sql" where the query is located. To run that command and create the necessary table, use the makefile statement by running the following command:

```bash
$ make initDB
```

### 3. Run test

In order to check if the project works, run the test command using the following makefile statement:

```bash
$ make testProject
```

You can check the project coverage using this other makefile command:

```bash
$ make coverage
```

## Starting the project

To start the project, run the following command. It starts the backend server on port :3000. (Ensure you are in the root project folder.)

```bash
$ go run main.go
```

```bash
URL = http://http://localhost:3000
```

### Endpoints

All these endpoints use the following URL:

#### POST - Create Task

```bash
URL = http://http://localhost:3000/tasks
```

```bash
Method = POST
```

This endpoint allows creating a task based on the information provided in the request body as shown in the following example. (Remember that the state can be any of the specific values: pending, in progress, completed.)

```json
{
  "title": "Task title",
  "description": "Task description",
  "state": "en progreso"
}
```

#### GET - Get Task

This endpoint allows getting a task based on the parameter (Id) provided in the URL as shown in the following example.

```bash
URL = http://http://localhost:3000/tasks/:id
```

```bash
Method = GET
```

#### PUT - Update Task

```bash
URL = http://http://localhost:3000/tasks
```

```bash
Method = PUT
```

This endpoint allows modifying a desired task based on the URL parameter. The program will update the task information with the data provided in the request body, as shown in the following example. (Remember that the state can be any of the specific values: pending, in progress, completed.)

```json
{
  "title": "Task title",
  "description": "Task description",
  "state": "en progreso"
}
```

#### DELETE - Delete Task

This endpoint allows deleting a task based on the parameter (Id) provided in the URL as shown in the following example.

```bash
URL = http://http://localhost:3000/tasks/:id
```

```bash
Method = DELETE
```

## Load Test

To assess the efficiency and scalability of the project, use load tests to determine how many requests the project can handle. Use this file to run the tests. You need to have Apache JMeter installed, at least version 5.6.3.

[Jmetter Load Test File](https://drive.google.com/file/d/10ub7CJDsAAWxoenb9NSzq99BAEW7tpZ6/view?usp=sharing)
