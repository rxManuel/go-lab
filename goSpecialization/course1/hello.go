package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	var a *int
	var b int

	// var a int
	// a = &b
	// cannot use &b (value of type *int) as int value in assignment
	a = &b
	b = 5
	fmt.Printf(" a: %d  *a: %d  b: %d  \n", a, *a, b)
	switch {
	case *a > 4: // each case implicit a break
		fmt.Println(" case 1")
	case b > 4:
		fmt.Println(" case 2")
	default:
		fmt.Println(" tagless switch")
	}

	switch {
	case *a > 4: // each case implicit a break BUT
		fmt.Println(" case 1")
		fallthrough //fallthrough
	case b > 4:
		fmt.Println(" case 2")
		fallthrough
	default:
		fmt.Println(" tagless switch")
	}
}
