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
	r.LoadHTMLGlob("views/*")

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

	r.GET("/items", func(c *gin.Context) {
		items := machine.ListItems()
		c.JSON(http.StatusOK, items)
	})

	r.POST("/buy", func(c *gin.Context) {
		name := c.Query("name")
		amount := machine.BuyItems(name)
		c.String(http.StatusOK, amount)

	})


	/* checkBalance */
	r.GET("/balance", func(c *gin.Context) {
		m := machine.GetWallet();
		if m == nil {
			c.String(http.StatusOK, "No Wallet")
			return
		} 
		c.String(http.StatusOK, strconv.Itoa( m.CheckBalance()))
	})
	/* insertCoin */
	r.GET("/balance/:amount", func(c *gin.Context) {
		m := machine.GetWallet();
		if m == nil {
			c.String(http.StatusOK, "No Wallet")
			return
		} 
		amount, err := strconv.Atoi(c.Param("amount"))
		if err != nil {
			c.String(http.StatusOK, "fail to parse")
			return
		}
		m.InsertCoin( amount )
		c.String(http.StatusOK, strconv.Itoa( m.CheckBalance()))
	})

	/* retriveCoin */
	r.GET("/retrive", func(c *gin.Context) {
		m := machine.GetWallet();
		if m == nil {
			c.String(http.StatusOK, "No Wallet")
			return
		} 
		m.RetriveCoin()
		c.String(http.StatusOK, strconv.Itoa( m.CheckBalance()))
	})

	/* vendingMachine INDEX */
	r.GET("/", func(c *gin.Context) {
		m := machine.GetWallet();
		c.HTML(http.StatusOK, "machineInterface.html", 
			gin.H{
				"title": "CSIM Vending Machine",
				"balance": m.CheckBalance(),
				"items": machine.ListItems(),
			},
		)
	})


	return r
}

func main() {
	// Init database
	machine.InitDb()
	// Init vendingMachine
	machine.NewWallet()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
