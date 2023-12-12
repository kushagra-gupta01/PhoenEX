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
	assert(t,len(ob.Asks),2)
}
