package main

/**
	RUN THIS PROGRAM MANY TIMES TO SEE THE UNEXPECTED RESULTS
	RACE CONDITION: Since there is no communitaction between go routines the value of global is not predictable
**/

import "fmt"

// global is a shared resource
var global = 0

func increase() {
	global++
	fmt.Println("Increasing shared global variable: ", global)
}

func decrease() {
	global--
	fmt.Println("Decreasing shared global variable: ", global)
}

func main() {
	fmt.Println("Initial value of Shared Global variable: ", global)
	go increase()
	go decrease()
	expected := 0

	for i := 0; i < 5; i++ {
		increase()
		decrease()
	}
	// Since increase and decrease are called n timeas each 0 value is expected
	if global == expected {
		fmt.Println("whoa! You got expected 0 very deterministic execution!")
	} else {
		fmt.Println("no what we expect so what is the value? the value is ", global)
	}

}
