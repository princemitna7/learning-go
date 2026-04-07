package main

import "fmt"

var b string = "Mitna"

func main() {
	a := "Prince"
	fmt.Printf("My name is %s %s\n", a, b)

	var c string
	fmt.Scanf("%s", &c)
	fmt.Println(c)

	var d bool
	fmt.Scan("%t", &d)
	// what does this do?

	const num = 100
}
