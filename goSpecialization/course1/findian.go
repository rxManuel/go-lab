package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var userString string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter an String: ")
	userString, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input...", err)
		return
	}

	userString = strings.TrimSpace(userString)

	if len(userString) > 2 { //ian 3 length
		if strings.HasPrefix(strings.ToLower(userString), "i") &&
			strings.Contains(strings.ToLower(userString), "a") &&
			strings.HasSuffix(strings.ToLower(userString), "n") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
