package main

import "fmt"

func modify(numbers ...int) {
        for i := range numbers {
                numbers[i] -= 5
        }
}

func main() {
	i := 10
	fmt.Printf("%T %v \n", &i, &i)
	fmt.Printf("%T %v \n", *(&i), *(&i))

	//declare a pointer

	var ptr_i *int = &i
	fmt.Println(ptr_i)

	var y int
    var ptr *int = &y
    fmt.Println(y)
    fmt.Println(*ptr)

	arr := []int{10, 20, 30}
    fmt.Println(arr)
    modify(arr...)
    fmt.Println(arr)
}
