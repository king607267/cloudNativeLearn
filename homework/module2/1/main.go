package main

import (
	"fmt"
	"time"
)

func main() {
	length := 10
	pipeline := make(chan int, length)
	done := make(chan bool)
	// producer1
	go func() {
		for i := 0; i < length; i++ {
			fmt.Printf("producer1 put value: %d\n", i)
			time.Sleep(time.Second)
			pipeline <- i

			if len(pipeline) == 10 {
				defer close(pipeline)
				close(done)
				return
			}
		}
	}()
	// producer2
	go func() {
		for i := 0; i < length; i++ {
			fmt.Printf("producer2 put value: %d\n", i)
			time.Sleep(time.Second)
			pipeline <- i
		}
		if len(pipeline) == 10 {
			defer close(pipeline)
			close(done)
			return
		}
	}()
	<-done
	// consumer
	for item := range pipeline {
		fmt.Printf("consumer get value: %d\n", item)
	}
}
