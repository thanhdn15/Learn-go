package main

import "fmt"

type Test struct {
	t string
}

func newDemo() TestInterface {
	var a TestInterface = Test{
		t: "strings",
	}

	return a
}

func (t Test) Demo() string {
	return "abc"
}

type Speaker interface {
	Speak() string
}

type Foo struct {
}

func (Foo) Speak() string {
	return "Hello, I'm foo"
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var someSpeaker Speaker = Foo{}
	fmt.Print(someSpeaker.Speak())

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "Hello"
	describe(i)

	product := make(map[string]interface{}, 0)

	product["name"] = "Iphone 13 Pro max"
	product["price"] = 31000000
	product["quantity"] = 40

	fmt.Print(product)

	// Ép kiểu dữ liệu Interface
	var m interface{} = "hello"

	s := m.(string)
	fmt.Println(s)

	s, ok := m.(string)
	fmt.Println(s, ok)

	f, ok := m.(float64)
	fmt.Println(f, ok)

	//f = m.(float64) // panic
	//fmt.Println(f)
}
