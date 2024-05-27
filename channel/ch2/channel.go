package main

import (
	"fmt"
	"strconv"
	"time"
)

func f1(ch chan string) {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		ch <- fmt.Sprintf("f1 %s", strconv.Itoa(i))
	}
	close(ch)
}

func f2(ch chan string) {
	for i := 0; i < 100; i++ {
		time.Sleep(2 * time.Second)
		ch <- fmt.Sprintf("f2 %s", strconv.Itoa(i))
	}
	close(ch)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go f1(ch1)
	go f2(ch2)
	//go f2(ch1, ch2)
	//for s := range ch2 {
	//fmt.Println(s)
	//}
	for i := 0; i < 100; i++ {
		select {
		case s := <-ch1:
			fmt.Println(s)
		case s := <-ch2:
			fmt.Println(s)
		}
	}
}
