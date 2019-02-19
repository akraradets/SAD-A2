package main

import (
	"github.com/SAD-A2/machine"
	// "github.com/SAD-A2/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
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

	// vendingMachine Singleton
	r.GET("/vending/name", func(c *gin.Context) {
		m := machine.GetMachine();
		if m == nil {
			c.String(http.StatusOK, "No Vending Machine")
			return
		} 
		result := m.Display()
		c.String(http.StatusOK, result)
	})

	r.GET("/vending/count", func(c *gin.Context) {
		m := machine.GetMachine();
		if m == nil {
			c.String(http.StatusOK, "No Vending Machine")
			return
		} 
		result := m.Count()
		c.String(http.StatusOK, strconv.Itoa(result))
	})

	return r
}

func main() {
	// Init database
	machine.InitDb()
	// Init vendingMachine
	machine.NewMachine("CSIM Machine")

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
