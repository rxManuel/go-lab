package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal Animal) Eat() {
	fmt.Println(animal)
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	cow := Animal{food: "grass", noise: "moo", locomotion: "walk"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hss"}

	animals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}

	fmt.Printf(" Ask me for an animal e.g. snake eat \n")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">")
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Invalid input, try 'cow speak' ")
			continue
		}

		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			fmt.Println("Enter valid data an animal and type (e.g. cow food) : NOT ", line)
			continue
		}
		animal, exist := animals[parts[0]]
		action := parts[1]

		if !exist {
			fmt.Println("Unknown animal, try again")
			continue
		}

		switch strings.TrimSpace(action) {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Unknown action, try: eat, move, speak")
		}

	}
}
