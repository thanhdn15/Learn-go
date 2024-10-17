package main

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

func main() {

}
