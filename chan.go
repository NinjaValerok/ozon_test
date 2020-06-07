package main

import (
	"fmt"
	"time"
)

type doSome func(int) int

func Merge2Channels(fn doSome, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {
	ff := func(inn1 <-chan int, inn2 <-chan int, outt chan<- int) {
		v1, v2 := <-inn1, <-inn2

		outt <- fn(v1) + fn(v2)
	}

	for i := 0; i < n; i++ {
		go ff(in1, in2, out)

	}

}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	fmt.Println("start long func")
	go Merge2Channels(longFunc, a, b, c, 3)
	fmt.Println("end long func")

	a <- 1
	b <- 2

	a <- 1
	b <- 2

	a <- 1
	b <- 2

	fmt.Println(<-c)
}

func longFunc(a int) int {
	fmt.Println("long func ", a)
	time.Sleep(1 * time.Second)
	return a + 1
}
