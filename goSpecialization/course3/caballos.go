package main

import (
	"fmt"
	"math/rand"
	"time"
)

func caballoCorre(id string, c chan string) {
	secs := rand.Intn(5) + 1
	time.Sleep(time.Duration(secs) * time.Second)
	c <- " El caballo " + id
}

func main() {
	caballo1 := make(chan string)
	caballo2 := make(chan string)
	caballo3 := make(chan string)

	go caballoCorre("blanco", caballo1)
	go caballoCorre("prieto asabache", caballo2)
	go caballoCorre("percheron", caballo3)

	fmt.Printf("El ganador es ")
	select {
	case ganador := <-caballo1:
		fmt.Println(ganador)
	case ganador := <-caballo2:
		fmt.Println(ganador)
	case ganador := <-caballo3:
		fmt.Println(ganador)
	}

	fmt.Printf("En segundo lugar llega ")
	select {
	case ganador := <-caballo1:
		fmt.Println(ganador)
	case ganador := <-caballo2:
		fmt.Println(ganador)
	case ganador := <-caballo3:
		fmt.Println(ganador)
	}

	fmt.Printf("El gran perdedor es ")
	select {
	case ganador := <-caballo1:
		fmt.Println(ganador)
	case ganador := <-caballo2:
		fmt.Println(ganador)
	case ganador := <-caballo3:
		fmt.Println(ganador)
	}
}
