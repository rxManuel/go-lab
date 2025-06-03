package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var nums []int = make([]int, 0, 3)
	var value string

	for i := 0; ; i++ {

		fmt.Print("Enter a number: ")
		fmt.Scan(&value)

		if strings.ToLower(value) == "x" {
			break
		} else {
			n, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Invalid input, please try again.")
				continue
			}
			// sorted insertion
			if i == 0 {
				nums = append(nums, n)
			} else {
				var position int
				// look for position
				for position = 0; position < len(nums); position++ {
					if n < nums[position] {
						break
					}
				}
				nums = append(nums, 0)                   // add space
				copy(nums[position+1:], nums[position:]) // move
				nums[position] = n
			}
		}
		fmt.Println(nums)
	}
}
