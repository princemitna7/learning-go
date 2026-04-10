package main

import "fmt"

func main() {
	fmt.Println("Operators")

	var a, b string = "Cat", "Dog"
	fmt.Println("Check if a equals to b:", a == b)

	var c string = "cat"
	var d int = 2
	// fmt.Println(c == d) this will give an error mismatched types
	// fmt.Println(c+d) error for mismatched types
	fmt.Printf("%T and %T\n", c, d)

	if a == "Cat" {
		fmt.Println("Yes its a cat")
	} else {
		fmt.Println("No")
	}

	switch b {
	case "Dog":
		fmt.Printf("Yes its Dog\n")
		fallthrough
	case "Any":
		fmt.Println("printanyway")
	}

	for i := 1; i <= 3; i++ {
		fmt.Println("Hello from loop")
	}

}
