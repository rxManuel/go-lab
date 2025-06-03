package main

import "fmt"

func main() {
	var myFloat float32
	fmt.Print("Please enter a floting number: ")
	fmt.Scan(&myFloat)

	fmt.Printf("%d \n", int(myFloat))
}
