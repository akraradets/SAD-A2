package main

import (
	"github.com/SAD-A2/machine"
	// "github.com/SAD-A2/controllers"
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"log"
	"github.com/carlescere/scheduler"
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

	// List All item in DB.
	r.GET("/items", func(c *gin.Context) {
		items := machine.ListItems()
		c.JSON(http.StatusOK, items)
	})


	r.POST("/pushButton/:name", func(c *gin.Context) {
		name := c.Params.ByName("name")
		pButton := machine.NewProxyButton(name)
		err := pButton.Push()
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
		} else {
			c.String(http.StatusOK, "Successed")
		}

	})


	/* Insert Coint */
	r.POST("/insertCoin", func(c *gin.Context) {
		amount, err := strconv.Atoi(c.PostForm("coin"))
		if err != nil {
			log.Panic(err)
			c.String(http.StatusOK, "fail to parse")
			return
		}
		m := machine.GetWallet();
		m.InsertCoin( amount )
		renderHTML(c)
	})

	/* retriveCoin */
	r.GET("/retrive", func(c *gin.Context) {
		m := machine.GetWallet();
		m.RetriveCoin()
		renderHTML(c)
	})

	/* vendingMachine INDEX */
	r.GET("/", func(c *gin.Context) {
		renderHTML(c)
	})

	return r
}

func renderHTML(c *gin.Context){
	m := machine.GetWallet();
	c.HTML(http.StatusOK,
		"machineInterface.html",
		gin.H{
			"balance": m.CheckBalance(),
			"items": machine.ListItems(),
		},
	)
}

func main() {
	// Init database
	machine.InitDb()
	// fetch items from database
	machine.LoadItems()
	// Init singleton Wallet
	machine.NewWallet()
	//scheduler to update database after every 5 minutes
	scheduler.Every(5).Minutes().Run(machine.UpdateItems)

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
