package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Person struct {
		fname string
		lname string
	}

	persons := make([]Person, 0, 3)
	var fileName string

	fmt.Print("What is the name of the file? ")
	fmt.Scan(&fileName)

	fileData, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error loading ", fileName)
		return
	}
	scanner := bufio.NewScanner(fileData)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.SplitN(line, " ", 2)

		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		fname := strings.TrimSpace(parts[0])
		lname := strings.TrimSpace(parts[1])

		persons = append(persons, Person{fname: fname, lname: lname})
	}
	for _, person := range persons {
		fmt.Printf("%s %s \n", person.fname, person.lname)
	}
}
