package main

import (
	"fmt"
	"time"
)

func mine(n int, transferPoint chan int){
	for i := 0; i <n; i++{
		fmt.Println("Miner %d is went to the mining center", i)
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Miner %d has mining", i)
		transferPoint <- i
	}

}


func main() {
	var coal int // this is for main goroutine to take from channel and take away the info needed
	transferPoint := make(chan int)//this is storage
	time.Since(time.Now())
	go mine(6, tranferPoint)
	for i :=0; i<6; i++{
		coal  <- transferPoint
	}
}
