/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and
how it can occur.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)


var counter int

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go incrementor("Foo:", &wg)
	go incrementor("Bar:", &wg)
	wg.Wait()
}

func incrementor(s string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		x := counter
		x++
		time.Sleep(10 * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

/*
- Race condition occurs when multithreaded code compete for a resource
and resulting final state depending on who gets the resource first.
- It's an anomalous behavior caused by the unexpected dependence on the relative timing of events.

- In this program a variable named counter of type int is declared as global variable.
- A function named incrementor will accept a string parameter for the naming purpose,
it will read counter variable and increment the value by 1, and prints out the current state value.
- Inside main function, two goroutines were created.

OUTPUT:
Bar: 0 Counter: 1
Foo: 0 Counter: 1
Foo: 1 Counter: 2
Bar: 1 Counter: 2
Foo: 2 Counter: 3
Bar: 2 Counter: 3
Bar: 3 Counter: 4
Foo: 3 Counter: 4
Bar: 4 Counter: 5
Foo: 4 Counter: 5

- The output showing that the program is having race condition.
- Running this code withs "go run -race race.go" will report "Found 1 data race"
*/