//package main
//
//import (
//	"fmt"
//	"runtime"
//)
//
//func main() {
//
//	var a int
//	a = test()
//
//	fmt.Print(a)
//	runtime.GC()
//}
//
//func test() int {
//	x := 1
//
//	return x
//}
//
//func test2() {
//	var x = new(int)
//
//	*x = 1
//	fmt.Println(x)
//}

import "fmt"

type Shape interface {
Area() float64
Perimeter() float64
}

type Rectangle struct {
Length, Width float64
}

func (r Rectangle) Area() float64 {
return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
return r.Length * r.Width
}

func main() {
var r Shape = Rectangle{Length: 3, Width: 4}
fmt.Printf("Type of r: %T, Area: %v, Perimeter: %v.", r, r.Area(), r.Perimeter())
}

