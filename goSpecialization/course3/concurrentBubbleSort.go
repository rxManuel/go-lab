package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func fillArray(array *[]int) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter integer numbers separated by space e.g. 1 2 3 5 8")
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Invalid input. Try: 4 3 1 2 8 6 5")
			continue
		}

		strNumbers := strings.Fields(line)
		if len(strNumbers) == 0 {
			fmt.Println("No valid input detected. Try: 5 6 8 4")
			continue
		}

		for _, s := range strNumbers {
			num, err := strconv.Atoi(s)
			if err == nil {
				*array = append(*array, num)
			} else {
				fmt.Printf("Invalid integer: %q\n", s)
			}
		}
		break
	}
}

func bubbleSort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return nums
}

func sortInChannel(array []int, c chan []int) {
	copySlice := append([]int(nil), array...)
	fmt.Println(copySlice)
	c <- bubbleSort(copySlice)
}

func main() {
	initialArray := make([]int, 0)
	fillArray(&initialArray)

	if len(initialArray) > 0 {

		if len(initialArray) > 3 {
			parts := 4
			jump := len(initialArray) / parts
			channel := make(chan []int, parts)
			go sortInChannel(initialArray[0:jump], channel)
			go sortInChannel(initialArray[jump:jump*2], channel)
			go sortInChannel(initialArray[jump*2:jump*3], channel)
			go sortInChannel(initialArray[jump*3:], channel)
			sorted1 := <-channel
			sorted1 = append(sorted1, (<-channel)...)
			sorted1 = append(sorted1, (<-channel)...)
			sorted1 = append(sorted1, (<-channel)...)
			fmt.Println(bubbleSort(sorted1))
		} else {
			fmt.Println(bubbleSort(initialArray))
		}

	} else {
		fmt.Println("Empty array, nothing to sort")
	}
}
