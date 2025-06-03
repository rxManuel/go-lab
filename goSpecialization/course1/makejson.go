package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")
	name, _ := reader.ReadString('\n')

	fmt.Print("What is you address? ")
	address, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)
	address = strings.TrimSpace(address)

	pmap := map[string]string{"name": name, "address": address}

	byteArr, err := json.Marshal(pmap)
	if err != nil {
		fmt.Println("Error Marshalling: ", pmap)
	} else {
		fmt.Println(string(byteArr))
	}

}
