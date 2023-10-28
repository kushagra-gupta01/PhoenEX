package main

import (
	"fmt"
	"time"
)

type Match struct{
	Asks 				*Order
	Bids 				*Order
	SizeFilled 	float64
	Price 			float64
}

type Order struct{
	Size 			float64
	Bid				bool
	Limit 		*Limit
	TimeStamp int64
}

func NewOrder(bid bool,size float64) *Order{
	return &Order{
		Size: size,
		Bid: bid,
		TimeStamp: time.Now().UnixNano(),
	}
}

func (o *Order)String() string{
	return fmt.Sprintf("[size:%.2f]",o.Size)
}

type Limit struct{
	Price 			 float64
	Orders			 []*Order
	TotalVolume  float64
}

func NewLimit(price float64) *Limit{
	return &Limit{
		Price: price,
		Orders: []*Order{},
	}
}

func (l *Limit)AddOrder(o *Order){
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (l *Limit)DeleteOrder(o *Order){
	for i:=0;i<len(l.Orders);i++{
		if l.Orders[i] == o{
			l.Orders[i] = l.Orders[len(l.Orders) - 1]
			l.Orders = l.Orders[:len(l.Orders) - 1]
		}
	}

	o.Limit = nil
	l.TotalVolume -= o.Size

	//TODO: resort the whole resting order
}

type OrderBook struct{
	Asks []*Limit
	Bids []*Limit

	AsksLimit map[float64]*Limit
	BidsLimit map[float64]*Limit
}

func NewOrderBook() *OrderBook{
	return &OrderBook{
		Asks: []*Limit{},
		Bids: []*Limit{},
		AsksLimit: make(map[float64]*Limit),
		BidsLimit: make(map[float64]*Limit),
	}
}

func (ob *OrderBook)PlaceOrder(price float64,o *Order) []Match{
	//1. try to match the orders with matching logic
	//2. add the rest of the order to the books
	if o.Size >0.0{
		ob.add(price, o)
	}
	return []Match{}
}

func (ob *OrderBook)add(price float64, o *Order){
	
}