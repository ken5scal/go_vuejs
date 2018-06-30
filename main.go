package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	_ "github.com/mattn/go-sqlite3"

	"database/sql"
	"github.com/ken5scal/go-vuejs/handlers"
)

// All the code is borrowed from https://scotch.io/tutorials/create-a-single-page-app-with-go-echo-and-vue
// This code serves as my practice for building go-backend and vue-frontend.
func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	// curl -i localhost:8080/tasks
	// curl -H 'Content-Type: application/json' -X PUT -d '{"name":"Foobar"}' localhost:8080/tasks
	// curl -i -X DELETE localhost:8080/tasks/1
	e.File("/", "public/index.html")
	e.Get("/tasks", handlers.GetTasks(db))
	e.Put("/tasks", handlers.PutTask(db))
	e.Delete("/tasks/:id", handlers.DeleteTask(db))
	//func(c echo.Context) error { return c.JSON(200, "Put Tasks")}


	e.Run(standard.New(":8080"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	} else if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS tasks(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name VARCHAR NOT NULL
    );
    `

	_, err := db.Exec(sql)
	// Exit if something goes wrong with our SQL statement above
	if err != nil {
		panic(err)
	}
}