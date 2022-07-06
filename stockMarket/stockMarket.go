package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumberGen(min float32, max float32) float32 {
	return min + (max-min)*rand.Float32()
}

type Stock struct {
	name  string
	price float32
}

var stockMarket = []Stock{
	{"GOOG", 2313.50},
	{"AAPL", 157.28},
	{"FB", 203.77},
	{"TWTR", 50},
}

func (stock *Stock) updateStock() {
	change := randomNumberGen(-50.00, 50.00)
	stock.price += change
}

func displayMarket(market []Stock) {
	fmt.Println("\nStock Prices as of", time.Now().Local().Format(time.RFC850))
	for _, stockName := range market {
		fmt.Printf("%v: %v\n", stockName.name, stockName.price)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	for {
		displayMarket(stockMarket)
		stockMarket[rand.Intn((3-0)+0)].updateStock()
		time.Sleep(5 * time.Second)
	}
}
