package main

import (
	"fmt"
	"time"
)
func main() {
	//channel:=make(chan type,buffer number)
	channel := make(chan int, 2)
	//goroutine
	go func() {
		time.Sleep(time.Second * 2)
		channel <- 8
	}() //Delay 2 seconds
	go func() { channel <- 9 }()

	//select to accomplish timeout elegantly
	for i:=0;i<2;i++{
	select {
	case msg := <-channel:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
		}
	}
}
