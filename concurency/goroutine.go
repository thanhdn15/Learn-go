package main

import (
	"fmt"
	"log"
	"sync"
)

// Create variable wg is type waitgroup
var wg = sync.WaitGroup{}

func printNumber() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d", i)
	}

	wg.Done()
}

func printChar() {
	for i := 'A'; i < 'A'+26; i++ {
		fmt.Printf("%c", i)
	}

	wg.Done()
}

func main() {

	// Declare number go routine run.
	wg.Add(2)

	// each goroutine, make a call Done before return
	go printChar()
	go printNumber()

	// call method wg.Wait()
	wg.Wait()

	log.Println("Main finished")
}
