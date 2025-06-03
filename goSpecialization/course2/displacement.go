package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return (0.5 * a * (t * t)) + (v0 * t) + s0
	}
}

func askData(label string, field *float64) {
	for {
		fmt.Printf("Please enter %s: ", label)
		_, err := fmt.Scanf("%f", field)
		if err != nil {
			fmt.Println("Invalid input, enter a floating number. Try again")
			continue
		} else {
			break
		}
	}
}

func main() {
	var ac, v0, s0, t float64
	askData("acceleration", &ac)
	askData("initial velocity", &v0)
	askData("initial displacement", &s0)
	fn := GenDisplaceFn(ac, v0, s0)
	askData("time", &t)
	fmt.Println(fn(t))
}
