package main

import (
	"log"
)

func printNumber(numChan chan int) {
	var result int

	for i := 0; i <= 100; i++ {
		result += i
	}

	numChan <- result
}

func printChar(strChan chan string) {
	var result string

	for i := 'A'; i < 'A'+26; i++ {
		result += string(i)
	}

	strChan <- result
}

func main() {

	chanPrintNumber := make(chan int)
	chanPrintChar := make(chan string)

	// each goroutine, make a call Done before return
	go printChar(chanPrintChar)
	go printNumber(chanPrintNumber)

	log.Println("Result Print Number: ", <-chanPrintNumber)
	log.Println("Result Print Char: ", <-chanPrintChar)

	log.Println("Main finished")
}
