package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createCounter(start int) chan int {
	channel := make(chan int)
	go func() {
		for {
			channel <- start
			start++
		}
	}()
	return channel
}

func DoubleCounter() {
	chanA := createCounter(10)
	chanB := createCounter(100)
	for i := 0; i < 10; i++ {
		fmt.Printf("CounterA=%d,CounterB=%d\n", <-chanA, <-chanB)
	}
	//defer close(chanA)
	//defer close(chanB)
}

func SelectChannel() {
	channels := make([]chan int, 6)
	for i := range channels{
		channels[i] = make(chan int)
	}
	go func() {
		for i:=0;i<10;i++{
			sleeps := time.Second * time.Duration(rand.Intn(10))
			fmt.Println("Will sleep", sleeps)
			time.Sleep(sleeps)
			index := rand.Intn(len(channels))
			channels[index] <- rand.Intn(10000)
		}
	}()
	for {
		select {
		case rd := <-channels[0]:
			fmt.Printf("Channel[0] generate value[%d]\n", rd)
		case rd := <-channels[1]:
			fmt.Printf("Channel[1] generate value[%d]\n", rd)
		case rd := <-channels[2]:
			fmt.Printf("Channel[2] generate value[%d]\n", rd)
		case rd := <-channels[3]:
			fmt.Printf("Channel[3] generate value[%d]\n", rd)
		case rd := <-channels[4]:
			fmt.Printf("Channel[4] generate value[%d]\n", rd)
		case rd := <-channels[5]:
			fmt.Printf("Channel[5] generate value[%d]\n", rd)
		}
	}
	time.Sleep(time.Second*1000)
}
