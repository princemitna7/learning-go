package main

import "fmt"

func main() {
	fmt.Println("Hello Arrays")

	var grades [5]int = [5]int{1, 2, 3, 4} //since it has 5 value last will be 0
	fmt.Println(grades)

	var fruits [3]string
	fmt.Println(fruits)

	r := [...]int{1, 2} // ellipses another way to init array
	fmt.Println(r)

	fmt.Println(len(grades)) // len function

	for i := 0; i < len(grades); i++ {
		fmt.Print(grades[i])
	}

	fmt.Print("\n")

	for index, element := range grades {
		fmt.Println(index, "=>", element) // using the range keyword
	}

	// 2d arrays
	arr := [2][2] int {{1,2},{3,4}}
	fmt.Println(arr)

	// slice

	slice := make([] int, 5, 8)

	fmt.Println(cap(slice)) //capacity function

	x := []int {-1,-2}

	for _, value:= range x{
		fmt.Println(value)
	}

	//slice append

	y := [5]int{10, 20, 90, 70, 60}
    s := y[:3]
    fmt.Println(cap(s))
    new_slice := append(s, 100, 200)
    fmt.Println(cap(new_slice))

	// append ...
	slice_2 := make([]int,5,20)
	new_slice_2 :=append(s,slice_2...)
	fmt.Println(cap(new_slice_2))

	fmt.Println(cap(slice_2))

	/*
	z := [5]int{10, 20, 90, 70, 60}
    p := append(z[:2], z[3:])  // since ... was not given it shows error
    fmt.Println(p)
	*/

	//maps

    ascii_codes := make(map[string]int)
    ascii_codes["A"] = 65
    ascii_codes["F"] = 70
    ascii_codes["K"] = 75
    fmt.Println(ascii_codes)

    ascii_codes = make(map[string]int)
    ascii_codes["U"] = 85
    fmt.Println(ascii_codes)

	my_arr :=[5]int{}
	fmt.Println(my_arr)

	fmt.Println(len(ascii_codes))

}
