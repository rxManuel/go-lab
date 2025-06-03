package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func calculateIndexes(size int) []int {
	result := make([]int, 4)
	i := 0
	for i < size {
		result[i%4] = result[i%4] + 1
		i++
	}
	return result
}

func goSort(arr []int, wg *sync.WaitGroup, c chan []int, piece int) []int {
	fmt.Printf("Subarray n%d length(%d): %d\n", piece, len(arr), arr)
	sortedArray := sort(arr)
	c <- sortedArray
	wg.Done()
	return sortedArray
}

func sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := sort(arr[:len(arr)/2])
	right := sort(arr[len(arr)/2:])
	return merge(left, right)
}

func merge(l []int, r []int) []int {
	mergedArr := []int{}
	i := 0
	j := 0
	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			mergedArr = append(mergedArr, l[i])
			i++
		} else {
			mergedArr = append(mergedArr, r[j])
			j++
		}
	}
	for i < len(l) {
		mergedArr = append(mergedArr, l[i])
		i++
	}
	for j < len(r) {
		mergedArr = append(mergedArr, r[j])
		j++
	}
	return mergedArr
}

func main() {
	fmt.Println("Please enter int values separated by space")
	var wg sync.WaitGroup
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	split := strings.Split(scanner.Text(), " ")
	if len(split) == 0 {
		log.Fatal("No values entered.")
	}
	var initialArray = []int{}
	for _, v := range split {
		j, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		initialArray = append(initialArray, j)
	}
	indexes := calculateIndexes(len(initialArray))
	c := make(chan []int, 4)
	wg.Add(4)
	go goSort(initialArray[:indexes[0]], &wg, c, 1)
	go goSort(initialArray[indexes[0]:indexes[0]+indexes[1]], &wg, c, 2)
	go goSort(initialArray[indexes[0]+indexes[1]:indexes[0]+indexes[1]+indexes[2]], &wg, c, 3)
	go goSort(initialArray[indexes[0]+indexes[1]+indexes[2]:], &wg, c, 4)
	arr1 := <-c
	arr2 := <-c
	arr3 := <-c
	arr4 := <-c
	finalMerge := merge(merge(merge(arr1, arr2), arr3), arr4)
	fmt.Printf("Sorted array length(%d): %d\n", len(finalMerge), finalMerge)
}
