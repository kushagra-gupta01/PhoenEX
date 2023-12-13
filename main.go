package main

import (
	"github.com/kushagra-gupta01/Cryto_Marketplace/orderbook"
	"github.com/labstack/echo/v4"
)

func main(){

	e := echo.New()
	ex := NewExchange()
	e.POST("/order",ex.handlePlaceOrder)
	e.Start(":3000")

}

type Market string

const(
	MarketETH Market = "ETH"
)

type Exchange struct{
	orderbooks map[Market]*orderbook.OrderBook
}

func NewExchange() *Exchange{
	return &Exchange{
		orderbooks: make(map[Market]*orderbook.OrderBook),
	}
}

func(ex *Exchange) handlePlaceOrder(c echo.Context) error{
	return c.JSON(200,"f4g")
}