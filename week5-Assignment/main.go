// main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Student struct
type Student struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Year  int     `json:"year"`
	GPA   float64 `json:"gpa"`
}

// In-memory database (ในโปรเจคจริงใช้ database)
var students = []Student{
	{ID: "1", Name: "John Doe", Email: "john@example.com", Year: 3, GPA: 3.50},
	{ID: "2", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 3.75},
	{ID: "3", Name: "Porawat Chawrainak", Email: "porawat@example.com", Year: 4, GPA: 2.98},
	{ID: "4", Name: "Jane Smith", Email: "jane@example.com", Year: 2, GPA: 4.00},
}

func getName(c *gin.Context) {
	NameQuery := c.Query("Name")

	if NameQuery != "" {
		filter := []Student{}
		for _, student := range students {
			if fmt.Sprint(student.Name) == NameQuery {
				filter = append(filter, student)

			}

		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, students)

}

func getEmail(c *gin.Context) {
	EmailQuery := c.Query("Email")

	if EmailQuery != "" {
		filter := []Student{}
		for _, student := range students {
			if fmt.Sprint(student.Email) == EmailQuery {
				filter = append(filter, student)

			}

		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, students)

}

func getYear(c *gin.Context) {
	YearQuery := c.Query("Year")

	if YearQuery != "" {
		filter := []Student{}
		for _, student := range students {
			if fmt.Sprint(student.Year) == YearQuery {
				filter = append(filter, student)

			}

		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, students)

}

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "healthy"})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/Name",getName)
		api.GET("/Email",getEmail)
		api.GET("/Year",getYear)
	}

	r.Run(":8080")

}
