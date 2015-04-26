package main

import (
	"fmt"
	"math/rand"
	"time"
	"strings"
)

func init(){
	fmt.Println("Channels.go file init function called")
}

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
	/**defer func() {
		close(chanA)
		close(chanB)
	}()*/
}

func SelectChannel() {
	channels := make([]chan int, 6)
	for i := range channels {
		channels[i] = make(chan int)
	}
	go func() {
		for {
			sleeps := time.Second * time.Duration(rand.Intn(10))
			fmt.Println("Will sleep", sleeps)
			time.Sleep(sleeps)
			index := rand.Intn(len(channels))
			channels[index] <- rand.Intn(10000)
		}
	}()
	var counter = 0
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
		default:
			counter++
			time.Sleep(1*time.Second)
		}
	}

}

func ErrorRecover(){
	var strType interface{} = "this is string"
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("Recover error",err)
		}
	}()
	intType := strType.(int)
	fmt.Println(intType)
}

func MakeAndAddSuffix(suffix string) (func(string)string){
	return func(name string)string{
		if !strings.HasSuffix(name,suffix){
			return name+suffix
		}
		return name
	}
}
