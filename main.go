package main

import (
	"github.com/SAD-A2/machine"
	"net/http"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
  // Init database
  machine.InitDb()
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

	r.GET("/machine_test", func(c *gin.Context) {
		m := machine.VendingMachine {
			Name: "Sam",
		}
		result := m.Display()
		c.String(http.StatusOK, result)
	})

	r.GET("/items", func(c *gin.Context) {
		items := machine.ListItems()
		c.JSON(http.StatusOK, items)
	})

	r.POST("/buy", func(c *gin.Context) {
		name := c.Query("name")
		amount := machine.BuyItems(name)
		c.String(http.StatusOK, amount)

	})
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
