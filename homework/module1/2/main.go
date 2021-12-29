package main

import (
	"fmt"
	"time"
)

func main() {
	length := 10
	pipeline := make(chan int, length)
	// producer
	go func() {
		for i := 0; i < length; i++ {
			fmt.Printf("producer1 put value: %d\n", i)
			time.Sleep(time.Second)
			pipeline <- i
		}
		defer close(pipeline)
	}()
	// consumer
	for item := range pipeline {
		fmt.Printf("consumer get value: %d\n", item)
	}
}
