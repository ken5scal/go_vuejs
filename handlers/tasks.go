package handlers

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"github.com/ken5scal/go-vuejs/models"
)

type H map[string]interface{}

func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var task models.Task

		// Map incoming JSON body to the new Task struct
		c.Bind(&task)

		// Add a task to db
		if id, err := models.PutTask(db, task.Name); err == nil {
			return c.JSON(http.StatusCreated, H {
				"created": id,
			})
		} else {
			return err
		}
	}
}

func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		if _, err := models.DeleteTask(db, id); err == nil {
			return c.JSON(http.StatusOK, H {
				"deleted": id,
			})
		} else {
			return err
		}

	}
}