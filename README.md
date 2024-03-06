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

Before manipulating the CRUD endpoints, notice there are two Auth endpoints, **/logIn** and **/signUp**.

### POST - SignUp (create user)

Allows to create a user using AWS Cognito. This endpoint stores a JWT cookie in the request

```bash
URL = http://localhost:3000/signUp
```

```bash
Method = POST
```

```json
{
  "name": "Andres",
  "password": "Hola#1228."
}
```

This is the body that must be sent, the password has to be:

- Minimum length: 8 characters
- Contains at least 1 number
- Contains at least 1 special character
- Contains at least 1 uppercase letter
- Contains at least 1 lowercase letter

### POST - LogIn (log in user)

Allows log in as a registered user. This endpoint stores a JWT cookie in the request

```bash
URL = http://localhost:3000/logIn
```

```bash
Method = POST
```

```json
{
  "name": "Andres",
  "password": "Hola#1228."
}
```

This is the body that must be sent, the password has to be:

- Minimum length: 8 characters
- Contains at least 1 number
- Contains at least 1 special character
- Contains at least 1 uppercase letter
- Contains at least 1 lowercase letter

### Active or deactivate the auth validation.

You can activate or deactivate the user validation middleware for specific endpoints, to do that, **change the boolean flag in the main.go file**

- true: The validation is activate
- false: The validation is deactivated

```go
router.POST("/tasks", middleware.ValidateUser(true, auth), controller.CreateTask) // User validation activated
router.GET("/tasks/:id", middleware.ValidateUser(false, auth), controller.GetTask) // User validation deactivated
```

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

## Postman Test

You can use this link for getting the Postman Collection in order to test the project
[Postman Collection](https://warped-station-765276.postman.co/workspace/Personal~f9053f70-c36a-4191-b7e3-d342fc1ad905/collection/20027571-79ea39a9-ada2-45b6-b576-e32a78788cf8?action=share&creator=20027571)

## Load Test

To assess the efficiency and scalability of the project, use load tests to determine how many requests the project can handle. Use this file to run the tests. You need to have Apache JMeter installed, at least version 5.6.3.

[Jmetter Load Test File](https://drive.google.com/file/d/10ub7CJDsAAWxoenb9NSzq99BAEW7tpZ6/view?usp=sharing)
