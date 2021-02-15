/*
Implement the dining philosopher’s problem with the following constraints/modifications.

1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5. The host allows no more than 2 philosophers to eat concurrently.
6. Each philosopher is numbered, 1 through 5.
7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a
line by itself, where <number> is the number of the philosopher.
8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a
line by itself, where <number> is the number of the philosopher.
 */

package main

import (
	"fmt"
	"sync"
)

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	id int
	leftCS, rightCS *ChopS
}

func (p Philo) eat(r chan *Philo, permission <-chan bool, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		r <- &p
		if <- permission {
			fmt.Println("starting to eat", p.id)
			p.leftCS.Lock()
			p.rightCS.Lock()

			fmt.Println("finishing eating", p.id)

			p.rightCS.Unlock()
			p.leftCS.Unlock()
		}
		<- r
	}
	wg.Done()
	fmt.Println("Philosopher", p.id, "is full")
}

func host(r <-chan *Philo, p chan<- bool) {
	for {
		if len(r) == 2 {
			p <- true
		}
	}
}

func main() {
	var wg sync.WaitGroup
	defer wg.Wait()

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{
			i+1,
			CSticks[i],
			CSticks[(i+1)%5],
		}
	}

	register := make(chan *Philo, 2)
	permission := make(chan bool)
	go host(register, permission)

	for i:=0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(register, permission, &wg)
	}
}
