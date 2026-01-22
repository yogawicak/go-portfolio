package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
1. Producer Consumer Pattern
*/
func producerConsumerPattern() {
	fmt.Println("Producer Consumer Pattern")

	// init context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// init channel
	ch := make(chan int)

	// init producer
	go func() {
		// loop data and send to channel
		for v := range 20 {
			ch <- v
		}
		close(ch)
	}()

	// init consumer
	go func() {
		// loop data and print
		for v := range ch {
			fmt.Println(v)
		}
	}()

	// wait until channel closed
	<-ctx.Done()
}

/*
2. Print 1-100 with 2 goroutines (Ping-Pong)
*/
func alternatePrint() {
	fmt.Println("\n=== 2. Alternate Print 1-20 ===")
	ping := make(chan bool)
	pong := make(chan bool)

	// init wait group
	var wg sync.WaitGroup

	// add 2 goroutines to wait group
	wg.Add(2)

	// Goroutine 1 - prints odd numbers
	go func() {
		defer wg.Done() // signal completion when goroutine finished
		for i := 1; i <= 100; i += 2 {
			<-ping // wait for signal
			fmt.Printf("Goroutine 1: %d\n", i)
			pong <- true // signal goroutine 2
		}
	}()

	// Goroutine 2 - prints even numbers
	go func() {
		defer wg.Done() // signal completion when goroutine finished
		for i := 2; i <= 100; i += 2 {
			<-pong // wait for signal
			fmt.Printf("Goroutine 2: %d\n", i)
			if i < 100 {
				ping <- true // signal goroutine 1
			}
		}
	}()

	ping <- true // start the sequence
	// time.Sleep(1 * time.Second)
	wg.Wait()
}

/*
3. Print ABCABC... (ring pattern)
*/
func ringPattern() {
	fmt.Println("\n=== 3. Ring Pattern ===")

	var wg sync.WaitGroup
	wg.Add(3)

	chA := make(chan bool)
	chB := make(chan bool)
	chC := make(chan bool)

	// Goroutine A
	go func() {
		defer wg.Done()
		for v := range 100 {
			<-chA
			fmt.Println("A", "Goroutine A", v)
			chB <- true
		}
	}()

	// Goroutine B
	go func() {
		defer wg.Done()
		for v := range 100 {
			<-chB
			fmt.Println("B", "Goroutine B", v)
			chC <- true
		}
	}()

	// Goroutine C
	go func() {
		defer wg.Done()
		for v := range 100 {
			<-chC
			fmt.Println("C", "Goroutine C", v)

			// if under 99, send signal to A
			if v < 99 {
				chA <- true
			}
		}
	}()

	// Start the ring
	chA <- true
	wg.Wait()
}

/*
4. Custom Worker Pool
*/
func workerPool() {
	fmt.Println("\n=== 4. Custom Worker Pool ===")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	var wg sync.WaitGroup

	numWorkers := 3

	// consume jobs
	for w := 0; w <= numWorkers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				fmt.Println("Worker", w, "processing job", j)
				results <- j * 2
			}
		}()
	}

	// send jobs
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}

		// close jobs after all jobs are sent
		close(jobs)
	}()

	// collect results
	for r := range results {
		fmt.Println("Result:", r)
	}

	wg.Wait()

}

/*
5.
*/

/*
6.
*/

func main() {
	// producerConsumerPattern()
	// alternatePrint()
	ringPattern()
}
