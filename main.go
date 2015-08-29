package main

import (
	"fmt"
	"runtime"
	"time"
)

func running(chtot chan int) {
	chtot <- counter()
}

func counter() int {
	total := 0
	for i := 0; i <= 10; i++ {
		total += i
		fmt.Println(i)
		time.Sleep(1 * time.Millisecond)
	}
	return total
}

func pingpong(in chan string, out chan string) {
	for i := 1; i < 11; i++ {
		s := <-in
		fmt.Println(s)
		if s == "Ping" {
			out <- "Pong"
		} else {
			out <- "Ping"
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// chtot := make(chan int)
	// go running(chtot)
	// total := counter()
	//
	// total += <-chtot
	//
	// time.Sleep(10 * time.Millisecond)
	// fmt.Println("end.", total)

	in := make(chan string, 1)
	out := make(chan string, 1)
	in <- "start"
	go pingpong(in, out)
	go pingpong(out, in)

	time.Sleep(10 * time.Second)
}
