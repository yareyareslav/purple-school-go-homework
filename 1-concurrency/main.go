package main

import (
	"fmt"
	"math/rand"
)

func createSlice(ch chan<- int) {
	go func() {
		defer close(ch)

		for range 10 {
			r := rand.Intn(100)
			ch <- r
		}
	}()
}

func square(ch <-chan int) <-chan int {
	res := make(chan int, 10)
	go func() {
		defer close(res)

		for a := range ch {
			res <- a * a
		}
	}()

	return res
} 

func main() {
	ch := make(chan int, 10)

	createSlice(ch)
	r := square(ch)

	for v := range r {
		fmt.Println(v)
	}
}