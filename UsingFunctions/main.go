package main

import "fmt"

func returnCube(n int)int{
	return n*n*n
}

func calcSquare(numbers []int) []int {
        squares := []int{}
        for _, v := range numbers {
                squares = append(squares, v*v)
        }
        return squares

}

func printStrings(s string, names ...string) {
        fmt.Println(s)
        for _, value := range names {
                fmt.Printf("%s, ", value)
        }
		fmt.Print("\n")

}

/*
func printString(names ...string) (names_c []string) {
    names_c = []string {}
        for _, value := range names {
                names_c = append(names_c, strings.ToUpper(value))
        }
        return
}
*/

//high order function

func addHundred(x int) int {
        return x + 100
}
func partialSum(x ...int) func() {
        sum := 0
        for _, value := range x {
                sum += value
        }
        return func() {
                fmt.Println(addHundred(sum))
        }
}

//defer statement

func main() {
	cube := returnCube(3)
	fmt.Printf("Cube of 3 is %d\n", cube)

	nums := [3]int{10,20,15}
	fmt.Println(calcSquare(nums[:]))

	printStrings("Hey there", "Joe", "Monica", "Gunther")

	fmt.Println("High order function")

	partial := partialSum(1,2,3,4,5)
	partial()
}
