package models

import "database/sql"

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

func GetTasks(db *sql.DB) TaskCollection {
	rows, err := db.Query("SELECT *FROM tasks")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		if err := rows.Scan(&task.ID, &task.Name); err != nil {
			panic(err)
		}
		result.Tasks = append(result.Tasks, task)
	}

	return result
}

func PutTask(db *sql.DB, name string) (int64, error) {
	// stmt means statement
	// a prepared statement can be compiled and cached
	// Prepared statements also prevent SQL Injection
	stmt, err := db.Prepare("INSERT INTO tasks(name) VALUES(?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		panic(err)
	}

	return result.LastInsertId()
}

func DeleteTask(db *sql.DB, id int) (int64, error) {
	// a prepared statement can be compiled and cached
	// Prepared statements also prevent SQL Injection
	stmt, err := db.Prepare("DELETE FROM tasks WHERE id = ?")
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	return result.RowsAffected()
}