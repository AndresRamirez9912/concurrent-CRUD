package models

import (
	"testing"

	"github.com/jmoiron/sqlx"
)

var task = Task{
	Id:          "2",
	Title:       "Test Task",
	Description: "Test Description",
	State:       "Test State",
}

func TestNewSqlConnection(t *testing.T) {
	repo := NewSqlConnection()
	defer func() {
		_ = repo.CloseConnection()
	}()

	t.Run("Open successful", func(t *testing.T) {
		err := repo.db.Ping()
		if err != nil {
			t.Errorf("Expected success connection, got %t", err)
		}
	})

}

func TestCloseConnection(t *testing.T) {
	repo := &SQLRepository{
		db: setupTestDatabase(t),
	}

	t.Run("Close successful", func(t *testing.T) {
		err := repo.CloseConnection()
		if err != nil {
			t.Error("Error closing the connection:", err)
		}
	})

}

func TestCreateTask(t *testing.T) {
	repo := &SQLRepository{
		db: setupTestDatabase(t),
	}
	defer repo.db.Close()

	t.Run("Success Creation", func(t *testing.T) {
		err := repo.CreateTask(task)
		if err != nil {
			t.Error("Error creating task:", err)
		}
	})
	t.Run("Error table doesn't exists", func(t *testing.T) {
		_, _ = repo.db.Exec(`DROP TABLE tasks`)
		err := repo.CreateTask(task)
		if err == nil {
			t.Error("Exec funtion must fail:", err)
		}
	})

	t.Run("Error beginning the transaction", func(t *testing.T) {
		repo.db.Close()
		err := repo.CreateTask(task)
		if err == nil {
			t.Error("Begin funtion must fail:", err)
		}
	})
}
func TestGetTask(t *testing.T) {
	repo := &SQLRepository{
		db: setupTestDatabase(t),
	}
	defer repo.db.Close()

	t.Run("Successfully gotten", func(t *testing.T) {
		myTask, err := repo.GetTask("1")
		if err != nil {
			t.Error("Error getting created task:", err)
		}

		if myTask.Title != task.Title {
			t.Errorf("Expected task title %s, got %s", task.Title, myTask.Title)
		}
	})

	t.Run("Error table doesn't exists", func(t *testing.T) {
		_, _ = repo.db.Exec(`DROP TABLE tasks`)
		_, err := repo.GetTask("1")
		if err == nil {
			t.Error("Exec funtion must fail:", err)
		}
	})

	t.Run("Error beginning the transaction", func(t *testing.T) {
		repo.db.Close()
		_, err := repo.GetTask("1")
		if err == nil {
			t.Error("Begin funtion must fail:", err)
		}
	})
}

func TestUpdateTask(t *testing.T) {
	repo := &SQLRepository{
		db: setupTestDatabase(t),
	}
	defer repo.db.Close()

	newTask := Task{
		Id:          "1",
		Title:       "New Test Task",
		Description: "New Test Description",
		State:       "New Test State",
	}

	t.Run("Successfull Update", func(t *testing.T) {
		err := repo.UpdateTask(newTask, task.Id)
		if err != nil {
			t.Error("Error updating the task:", err)
		}
	})

	t.Run("No existing id", func(t *testing.T) {
		err := repo.UpdateTask(newTask, "2")
		if err != nil {
			t.Error("Error updating the task:", err)
		}

		myTask, _ := repo.GetTask(task.Id)
		if myTask.Title != "" {
			t.Errorf("Expected task title empty, got %s", myTask.State)
		}
	})

	t.Run("Error table doesn't exists", func(t *testing.T) {
		_, _ = repo.db.Exec(`DROP TABLE tasks`)
		err := repo.UpdateTask(newTask, "2")
		if err == nil {
			t.Error("Exec funtion must fail:", err)
		}
	})

	t.Run("Error beginning the transaction", func(t *testing.T) {
		repo.db.Close()
		err := repo.UpdateTask(newTask, "2")
		if err == nil {
			t.Error("Begin funtion must fail:", err)
		}
	})

}

func TestDeleteTask(t *testing.T) {
	repo := &SQLRepository{
		db: setupTestDatabase(t),
	}
	defer repo.db.Close()

	t.Run("Successfull Update", func(t *testing.T) {
		err := repo.DeleteTask(task.Id)
		if err != nil {
			t.Error("Error deleting the task:", err)
		}
	})

	t.Run("Error table doesn't exists", func(t *testing.T) {
		_, _ = repo.db.Exec(`DROP TABLE tasks`)
		err := repo.DeleteTask("2")
		if err == nil {
			t.Error("Exec funtion must fail:", err)
		}
	})

	t.Run("Error beginning the transaction", func(t *testing.T) {
		repo.db.Close()
		err := repo.DeleteTask("2")
		if err == nil {
			t.Error("Begin funtion must fail:", err)
		}
	})
}

func setupTestDatabase(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("Error creating test database: ", err)
	}

	_, err = db.Exec(`
		CREATE TABLE tasks (
			id INTEGER PRIMARY KEY,
			title TEXT,
			description TEXT,
			state TEXT
		)
	`)

	if err != nil {
		t.Fatalf("Error creating tasks table: %v", err)
	}

	_, err = db.Exec(`INSERT INTO tasks (id, title, description, state) VALUES (1, 'Test Task', 'Test Description', 'Test State')`)

	if err != nil {
		t.Fatalf("Error creating the first task: %v", err)
	}

	return db
}
