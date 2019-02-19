package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

  // Ping test
	r.GET("/db", func(c *gin.Context) {
    result, err := test_connection()
    if err != nil {
      panic(err)
    }
		c.String(http.StatusOK, result)
	})

	r.GET("/machine_test", func(c *gin.Context) {
		m := vendingMachine.VendingMachine {
			Name: "Sam",
		}
		result := m.Display()
		c.String(http.StatusOK, result)
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
