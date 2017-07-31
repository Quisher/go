package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func produce(ch chan string) {
	ch <- "hello"
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	ch := make(chan string)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	go produce(ch)
	x, y := <-c, <-c // receive from c

	str := <-ch
	fmt.Println(str)
	fmt.Println(x, y, x+y)
}
