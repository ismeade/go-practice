package main

import (
	"fmt"
	"time"
)

func f1(ch chan<- int) {
	time.Sleep(1000 * time.Millisecond)
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 <-chan int, ch2 chan int) {
	for {
		val, ok := <-ch1
		if !ok {
			fmt.Println("break")
			break
		}
		ch2 <- val * val
	}

	//for key := range ch1 {
	//	ch2 <- key * key
	//}
	close(ch2)
}

func main() {

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)

	go f2(ch1, ch2)

	for key := range ch2 {
		fmt.Println(key)
	}
}
