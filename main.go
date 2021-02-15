package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type courses struct {
	id   int    `json:"id"`
	name string `json:"name"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Groups
	courses := e.Group("/course")
	//e.GET("/course", )
	courses.POST("", SetCourse)
	//e.DELETE("/course", )
	//e.PUT("/course", )

	e.Logger.Fatal(e.Start(":8080"))

}

// SetCourse .....
func SetCourse(c echo.Context) error {
	var data courses
	err := c.Bind(&data)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%+v", data)
	return c.JSON(http.StatusOK, map[string]string{"message": "Ok"})
}
