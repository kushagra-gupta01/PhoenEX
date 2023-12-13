package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T,a,b any){
	if !reflect.DeepEqual(a,b){
		t.Errorf("%+v != %+v",a,b)
	}
}

func TestLimit(t *testing.T){
	l := NewLimit(10_000)
	buyOrder := NewOrder(true,5)

	l.AddOrder(buyOrder)
	fmt.Println(l)
}

func TestPlaceLimitOrder(t *testing.T){
	ob := NewOrderBook()

	sellOrderA := NewOrder(false,10)
	sellOrderB := NewOrder(false,5)
	ob.PlaceLimitOrder(10_000,sellOrderA)
	ob.PlaceLimitOrder(9_000,sellOrderB)
	assert(t,len(ob.asks),2)
}

func TestPlaceMarketOrder(t *testing.T){
	ob:= NewOrderBook()

	sellOrder := NewOrder(false,20)
	ob.PlaceLimitOrder(10_000,sellOrder)

	buyOrder := NewOrder(true,10)
	matches := ob.PlaceMarketOrder(buyOrder)

	assert(t,len(matches),1)
	assert(t,len(ob.asks),1)
	assert(t,ob.AskTotalVolume(),10.0)
	
}

func TestPlaceMarketOrderMultiFill(t *testing.T){

}