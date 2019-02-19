package main

import (
	"github.com/SAD-A2/machine"
	"github.com/SAD-A2/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

  // Ping test
	r.GET("/db_test", func(c *gin.Context) {
    if err := machine.DB.Ping(); err != nil {
      c.String(http.StatusOK, "Failed")
    } else {
  		c.String(http.StatusOK, "Successed")
    }
	})

	// Machine Test
	r.GET("/machine_test", func(c *gin.Context) {
		m := machine.VendingMachine {
			Name: "Sam",
		}
		result := m.Display()
		c.String(http.StatusOK, result)
	})

	// Controller Test
	r.GET("/vendingMachine/index", func(c *gin.Context) {
		controller := controllers.VendingMachine{}
		result := controller.Index()
		c.String(http.StatusOK, result)
	})

	// Controller Test
	r.GET("/vendingMachine/name", func(c *gin.Context) {
		m := machine.VendingMachine {
			Name: "Sam Ja",
		}
		controller := controllers.VendingMachine{
			Machine: m,
		}
		result := controller.Name()
		c.String(http.StatusOK, result)
	})

	return r
}

func main() {
	// Init database
  machine.InitDb()


	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
