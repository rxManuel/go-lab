package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id              int
	feedCounter     int
	leftCs, rightCs *ChopS
}

type Host chan *Philo

func (f *Philo) eat(host Host, chopstickChannel chan *ChopS) {
	for f.feedCounter > 0 {
		host <- f                      //block in case of 2 philosofers in channel
		f.leftCs = <-chopstickChannel  // choose chopstick as avialble in channel
		f.rightCs = <-chopstickChannel // choose chopstick as avialble in channel

		f.leftCs.Lock()
		f.rightCs.Lock()

		fmt.Println("starting to eat ", f.id)
		f.feedCounter--
		fmt.Println("finishing eating", f.id)

		f.rightCs.Unlock()
		f.leftCs.Unlock()

		chopstickChannel <- f.leftCs // put chopstick back into channel
		chopstickChannel <- f.rightCs
		<-host // release space in host channel
	}
}

func invitePhilosopherToEat(philo *Philo, host Host, chopstickChannel chan *ChopS, wg *sync.WaitGroup) {
	defer wg.Done()
	philo.eat(host, chopstickChannel)
}

func main() {
	const nPhilos = 5
	const maxEating = 2 // max of philosopher to be hosted

	CSticks := make([]*ChopS, nPhilos)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	// Chopstick channel for req: The philosophers pick up the chopsticks in any order, not lowest-numbered first
	chopstickChannel := make(chan *ChopS, nPhilos)
	for _, chopstick := range CSticks {
		chopstickChannel <- chopstick
	}

	filos := make([]*Philo, nPhilos)
	for i := 0; i < nPhilos; i++ {
		filos[i] = &Philo{id: i, feedCounter: 3}
	}

	var wg sync.WaitGroup
	host := make(chan *Philo, maxEating)

	for _, f := range filos {
		wg.Add(1)
		go invitePhilosopherToEat(f, host, chopstickChannel, &wg)
	}
	wg.Wait()
}
