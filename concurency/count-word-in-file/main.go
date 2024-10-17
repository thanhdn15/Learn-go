package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func countFirstFile(result chan int, filePath string, keyword string) {
	var numberOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
		result <- 0
		return
	}
	numberOfOcc = strings.Count(string(fileContent), keyword)

	result <- numberOfOcc
	defer close(result)
}

func countSecondFIle(result chan int, filePath string, keyword string) {
	var numberOfOcc int
	fileContent, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Println(err)
		result <- 0
		return
	}

	numberOfOcc = strings.Count(string(fileContent), keyword)

	result <- numberOfOcc
	defer close(result)
}

func main() {
	countFirstChan := make(chan int)
	countSecondChan := make(chan int)

	go countFirstFile(countFirstChan, "1.txt", "bạn")
	go countSecondFIle(countSecondChan, "2.txt", "bạn")

	fmt.Print(<-countFirstChan)
	fmt.Print(<-countSecondChan)
}
