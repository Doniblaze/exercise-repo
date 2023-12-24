package section2

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func RunChannelExercise() {
	// Producer-Consumer scenario
	numbersCh := make(chan int)
	doneCh := make(chan struct{})

	// Producer goroutine
	go produceNumbers(numbersCh, doneCh)

	// Consumer goroutine
	go consumeAndPrintSquares(numbersCh, doneCh)

	// Wait for the goroutines to finish
	time.Sleep(2 * time.Second)
	close(doneCh)
}
func produceNumbers(numbersCh chan<- int, doneCh <-chan struct{}) {
	defer close(numbersCh)

	rand.Seed(time.Now().UnixNano())

	for {
		select {
		case <-doneCh:
			return
		default:
			// Generate a random number and send it to the channel
			num := rand.Intn(100)
			numbersCh <- num

			// Sleep for a short duration
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func consumeAndPrintSquares(numbersCh chan int, doneCh <-chan struct{}) {
	var wg sync.WaitGroup
	defer wg.Wait()

	for {
		select {
		case num, ok := <-numbersCh:
			if !ok {
				// Channel closed, exit the goroutine
				return
			}

			wg.Add(1)
			// Goroutine to calculate square and print
			go func(n int) {
				defer wg.Done()
				square := n * n
				fmt.Printf("Received: %d, Square: %d\n", n, square)
			}(num)

		case <-doneCh:
			// Close numbersCh when doneCh is closed
			close(numbersCh)
			return
		}
	}
}
