package main

import "time"

type Order struct{
	Size 			float64
	Bid				bool
	Limit 		*Limit
	TimeStamp int64
}

func NewOrder(size float64,bid bool) *Order{
	return &Order{
		Size: size,
		Bid: bid,
		TimeStamp: time.Now().UnixNano(),
	}
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

type OrderBook struct{
	Asks []*Limit
	Bids []*Limit
}