package main

import "fmt"

func demo1(abc chan int) {
	var i = 1

	abc <- i

}

func demo2(bcd chan int) {
	var i = 2

	bcd <- i
}

func main() {
	var d = make(chan int)
	var f = make(chan int)
	go demo1(d)
	go demo2(f)

	fmt.Print(<-d)
	fmt.Print(<-f)
}
