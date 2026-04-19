package main

import "fmt"

type Student struct {
	name   string
	rollNo int
	marks  []int
	grades map[string]int
}

type movie struct {
	name   string
	rating float32
}

type Rectangle struct {
	length  int
	breadth int
}

func (r Rectangle) area() int {
	return r.length * r.breadth
}

func (r *Rectangle) incLength(n int) {
	for i := 0; i < n; i++ {
		r.length += i
	}
}

type myall interface {
	area()
	incLength()
}

func main() {
	var s Student
	fmt.Printf("%+v", s)

	m := movie{name: "ABCD"}
	var m2 movie
	fmt.Printf("%+v", m)
	fmt.Printf("%+v", m2)

	r := Rectangle{breadth: 10, length: 5}
	fmt.Println(r.area())
	fmt.Println(r)
	r.incLength(7)
	fmt.Println(r.area())
	fmt.Println(r)
}
