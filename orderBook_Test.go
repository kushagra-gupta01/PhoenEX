package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T){
	l := NewLimit(10_000)
	buyOrder := NewOrder(true,5)

	l.AddOrder(buyOrder)
	fmt.Println(l)
}

func TestOrderBook(t *testing.T){
	ob := NewOrderBook()

	buyOrderA := NewOrder(true,13)
	buyOrderB := NewOrder(true,300)

	ob.PlaceOrder(18_000,buyOrderA)
	ob.PlaceOrder(9_000,buyOrderB)

	fmt.Printf("%+v",ob.Bids[0])
}
