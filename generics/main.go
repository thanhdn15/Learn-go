package main

import "fmt"

func Print[T any](s ...T) {
	for _, v := range s {
		fmt.Print(v)
	}
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Swap[T any](a, b T) (T, T) {
	return b, a
}

// Khai báo một interface ràng buộc kiểu dữ liệu số
type Odered interface {
	int | float64
}

func Max[T Odered](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

// Struct
type Container[T any] struct {
	value T
}

func (c *Container[T]) SetValue(v T) {
	c.value = v
}

func (c *Container[T]) GetValue() T {
	return c.value
}

func Map[T any, R any](s []T, f func(T) R) []R {
	result := make([]R, len(s))

	for i, v := range s {
		result[i] = f(v)
	}

	return result
}

func main() {
	Print("Hello, ", "playgound\n")

	//fmt.Print(MaxInt(2, 5))
	fmt.Print(Max(2, 1))
	x, y := Swap(10, 20)

	fmt.Print(x, y)

	// struct
	intContainer := Container[int]{value: 100}
	fmt.Println(intContainer.GetValue())

	stringContainer := Container[string]{value: "ABC"}
	fmt.Println(stringContainer.GetValue())

	nums := []int{1, 2, 3, 4}
	result := Map(nums, func(n int) int {
		return n * n
	})

	fmt.Print(result)
}
