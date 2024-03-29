package models

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Repository interface {
	CreateTask(Task) error
	GetTask(string) (*Task, error)
	UpdateTask(Task, string) error
	DeleteTask(string) error

	SignUpUser(*User) error
	LogInUser(string) (*User, error)
}

type SQLRepository struct {
	db *sqlx.DB
}

func NewSqlConnection() *SQLRepository {
	db, err := sqlx.Open("sqlite3", "tasks.db")
	if err != nil {
		log.Fatal("Error trying to connect to the DB", err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to the DB", err)
	}

	return &SQLRepository{db: db}
}

func (repo *SQLRepository) CloseConnection() error {
	err := repo.db.Close()
	if err != nil {
		log.Fatal("Error closing the connection")
		return err
	}
	return nil
}

func (repo *SQLRepository) CreateTask(task Task) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.Exec("INSERT INTO tasks (title, description, state) VALUES ($1,$2,$3)", task.Title, task.Description, task.State)
	if err != nil {
		log.Println("Error Trying to create task", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLRepository) GetTask(id string) (*Task, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	row, err := tx.Query("SELECT * FROM tasks WHERE id=$1", id)
	if err != nil {
		log.Println("Error trying to get the task", err)
		return nil, err
	}
	defer func() {
		err := row.Close()
		if err != nil {
			log.Fatal("Error trying to close the query", err)
		}
	}()

	task := Task{}
	for row.Next() {
		err = row.Scan(&task.Id, &task.Title, &task.Description, &task.State)
		if err != nil {
			log.Println("Error getting the information ", err)
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (repo *SQLRepository) UpdateTask(task Task, id string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.Exec("UPDATE tasks SET title=$1, description=$2, state=$3 WHERE id=$4", task.Title, task.Description, task.State, id)
	if err != nil {
		log.Println("Error Trying to update task", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLRepository) DeleteTask(id string) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		log.Println("Error Trying to delete task", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLRepository) SignUpUser(user *User) error {
	tx, err := repo.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	_, err = tx.Exec("INSERT INTO users (name, password) VALUES ($1,$2)", user.Name, user.Password)
	if err != nil {
		log.Println("Error Trying to create user", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (repo *SQLRepository) LogInUser(userName string) (*User, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	row, err := tx.Query("SELECT * FROM users WHERE name=$1", userName)
	if err != nil {
		log.Println("Error Trying to get the user info", err)
		return nil, err
	}

	defer func() {
		err := row.Close()
		if err != nil {
			log.Fatal("Error trying to close the query", err)
		}
	}()

	newUser := User{}
	for row.Next() {
		err = row.Scan(&newUser.Name, &newUser.Password)
		if err != nil {
			log.Println("Error getting the information ", err)
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}
