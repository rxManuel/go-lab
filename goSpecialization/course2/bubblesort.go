package main

import "fmt"

func fillSlice(nums []int, capacity int) {
	for i := 0; i < capacity; i++ {
		fmt.Printf("Enter a int number for element [%d] : ", i)
		_, err := fmt.Scanf("%d", &nums[i])
		if err != nil {
			fmt.Println("Invalid input, try again")
			i--
		}
	}
}

func swap(nums []int, posi int) {
	temp := nums[posi+1]
	nums[posi+1] = nums[posi]
	nums[posi] = temp
}

func BubbleSort(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1; j++ {
			if nums[j] > nums[j+1] {
				swap(nums, j)
			}
		}
	}
}

func main() {
	const capacity = 10
	elements := make([]int, capacity)
	fillSlice(elements, capacity)
	fmt.Println("Before sort: \n", elements)
	BubbleSort(elements)
	fmt.Println("After sort: \n", elements)
}
