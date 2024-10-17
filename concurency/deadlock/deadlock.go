package main

import "fmt"

func main() {
	chanInt := make(chan int)

	go handle(chanInt)

	fmt.Print(<-chanInt)
}

func handle(chanInt chan int) {
	i := 5
	chanInt <- i
}
