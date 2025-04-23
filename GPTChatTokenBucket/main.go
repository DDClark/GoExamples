package main

import (
	"fmt"
	"math/rand"
	"time"
)

const tokensPerInterval = 100
const interval = time.Second

var tokens chan struct{}

func init() {
	tokens = make(chan struct{}, tokensPerInterval)
	for i := 0; i < tokensPerInterval; i++ {
		tokens <- struct{}{}
	}
	go func() {
		for {
			time.Sleep(interval)
			tokens <- struct{}{}
		}
	}()
}

func rateLimit() bool {
	select {
	case <-tokens:
		return true
	default:
		return false
	}
}

func main() {
	for i := 0; i < 200; i++ {
		timeToWait := rand.Intn(100)
		time.Sleep(time.Duration(timeToWait) * time.Millisecond)
		if rateLimit() {
			fmt.Println("Request processed")
		} else {
			fmt.Println("Request rate limited")
		}
	}
}
