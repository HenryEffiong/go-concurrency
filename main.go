// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func printSomething(s string, wg *sync.WaitGroup) {
// 	defer wg.Done() // decrement wg by one after this function completes

// 	fmt.Println(s)
// }

// func main() {
// 	// create a variable of type sync.WaitGroup
// 	var wg sync.WaitGroup

// 	// this slice consists of the words we want to print using a goroutine
// 	words := []string{
// 		"alpha",
// 		"beta",
// 		"delta",
// 		"gamma",
// 		"pi",
// 		"zeta",
// 		"eta",
// 		"theta",
// 		"epsilon",
// 	}

// 	// we add the length of our slice (9) to the waitgroup
// 	wg.Add(len(words))

// 	for i, x := range words {
// 		// call printSomething as a goroutine, and hand it a pointer to our
// 		// waitgroup, since you never want to copy a waitgroup after it has
// 		// been created, or bad things happen...
// 		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
// 	}

// 	// our program will pause at this point, until wg is 0
// 	wg.Wait()

// 	// we have to add one to wg or we'll get an error when we call
// 	// printSomething again, since wg is already at 0
// 	wg.Add(1)
// 	printSomething("This is the second thing to be printed!", &wg)
// }

package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	// challenge: modify this code so that the calls to updateMessage() on lines
	// 27, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),
	// printMessage(), and main().

	msg = "Hello, world!"

	wg.Add(1)
	go updateMessage("Hello, universe!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, cosmos!")
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello, world!")
	wg.Wait()
	printMessage()
}

