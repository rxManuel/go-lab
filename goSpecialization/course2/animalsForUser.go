package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Animals struct {
	food       string
	locomotion string
	noise      string
}

type UserAnimal struct {
	name       string
	animalType Animals
}

var (
	Cow   = Animals{food: "grass", noise: "moo", locomotion: "walk"}
	Bird  = Animals{food: "worms", locomotion: "fly", noise: "peep"}
	Snake = Animals{food: "mice", locomotion: "slither", noise: "hss"}
)

func (animal Animals) Eat() {
	fmt.Println(animal.food)
}

func (animal Animals) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animals) Speak() {
	fmt.Println(animal.noise)
}

func findAnimal(animals []UserAnimal, names string) (Animals, error) {
	found := false
	idx := 0
	for i := 0; i < len(animals); i++ {
		if animals[i].name == names {
			found = true
			idx = i
			break
		}
	}
	if !found {
		fmt.Println(names, " was not found, try again.")
		return Animals{}, errors.New(" Animal not found")
	} else {
		return animals[idx].animalType, nil
	}
}

func newUserAnimal(animals *[]UserAnimal, names, animalType string) {
	switch animalType {
	case "cow":
		*animals = append(*animals, UserAnimal{name: names, animalType: Cow})
	case "snake":
		*animals = append(*animals, UserAnimal{name: names, animalType: Snake})
	case "bird":
		*animals = append(*animals, UserAnimal{name: names, animalType: Bird})
	default:
		fmt.Println(animalType, " is an invalid animal. Try: bird, cow or snake")
	}

	fmt.Println(animals)
}

func queryUserAnimal(animals []UserAnimal, names, action string) {
	var animal Animals
	animal, err := findAnimal(animals, names)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println(action, " is an invalid action. Try: eat, move or speak")
	}
	fmt.Println(animals)
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	userAnimals := make([]UserAnimal, 0, 2)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("Invalid input, try 'newanimal Lola cow' ")
			continue
		}

		parts := strings.Split(line, " ")

		if len(parts) != 3 {
			fmt.Println("Enter valid data an animal and type (e.g. cow food) : NOT ", line)
			continue
		}

		action := parts[0]
		switch strings.TrimSpace(action) {
		case "newanimal":
			newUserAnimal(&userAnimals, strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2]))
		case "query":
			queryUserAnimal(userAnimals, strings.TrimSpace(parts[1]), strings.TrimSpace(parts[2]))
		default:
			fmt.Println(action, " is an invalid action, try: newanimal or query")
		}

	}
}
