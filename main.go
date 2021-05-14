package main

import (
	"fmt"
	blockchain "mycoin/part6_2"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", hello)
	e.GET("/createwallet", createWallet)
	e.GET("/getbalance/:pubkey", getBalance)
	e.GET("/histories-all", historiesAll)
	e.POST("/mining", mining)
	e.POST("/send", sendCoin)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createWallet(c echo.Context) error {
	address := blockchain.CreateWallet()
	var response struct {
		Address string
	}
	response.Address = address
	return c.JSON(http.StatusOK, response)
}

func getBalance(c echo.Context) error {
	pubkey := c.Param("pubkey")
	balance := blockchain.GetBalance(pubkey)
	var response struct {
		Balance int
	}
	response.Balance = balance
	return c.JSON(http.StatusOK, response)
}

func historiesAll(c echo.Context) error {
	response := blockchain.PrintChain()
	return c.JSON(http.StatusOK, response)
}

func sendCoin(c echo.Context) error {
	var json struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int    `json:"amount"`
	}

	if err := c.Bind(&json); err != nil {
		return err
	}
	var response struct {
		Result string
	}
	response.Result = blockchain.Send(json.From, json.To, json.Amount)
	fmt.Println("result", response.Result)
	return c.JSON(http.StatusOK, response)
}

func mining(c echo.Context) error {
	var json struct {
		Address string `json:"address"`
	}

	if err := c.Bind(&json); err != nil {
		return err
	}
	var response struct {
		Result string
	}
	response.Result = blockchain.InitBlockchain(json.Address)
	fmt.Println("result", response.Result)
	return c.JSON(http.StatusOK, response)
}
