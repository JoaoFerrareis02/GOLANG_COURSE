package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	// close(doneChan) // Nessesary for 2º way. This will only work if you know that this is the last one to do
}

func main() {

	// 1. Running Functions as Goroutines
	// go greet("Nice to meet you!")
	// go greet("How are you?")
	// go slowGreet("How ... are ... you ...?")
	// go greet("I hope you're liking the course!")

	// 2. Introducing & Using channels
	// done := make(chan bool)
	// go slowGreet("How ... are ... you ...?", done)
	// <-done

	// 3. Working with Multiple Channels & Goroutines

	// 1º way

	dones := make([]chan bool, 4)

	for index := range dones {
		dones[index] = make(chan bool)
	}

	go greet("Nice to meet you!", dones[0])
	go greet("How are you?", dones[1])
	go slowGreet("How ... are ... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])

	for _, done := range dones {
		<-done
	}

	// 2º way

	// done := make(chan bool)

	// go greet("Nice to meet you!", done)
	// go greet("How are you?", done)
	// go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!", done)

	// for range done {
	// }

}
