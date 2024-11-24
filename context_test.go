package latihangolangcontext

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

// background := context

func TestContext(t *testing.T) {

	// membuat context
	backgronud := context.Background()
	fmt.Println(backgronud)

	todo := context.TODO()
	fmt.Println(todo)

	// parent dan child

}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "d")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)

	fmt.Println(contextF.Value("f"))
	fmt.Println(contextF.Value("c"))
	fmt.Println(contextF.Value("a"))

}

func CreateCounter(ctx context.Context, wg *sync.WaitGroup) chan int {

	destination := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(destination)
		counter := 1

		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++

				time.Sleep(1 * time.Second)
			}
		}

	}()

	return destination

}

// func TestContextWithCancel(t *testing.T) {
// 	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

// 	parent := context.Background()

// 	ctx, cancel := context.WithCancel(parent)

// 	destination := CreateCounter(ctx)

// 	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

// 	for i := range destination {
// 		fmt.Println("Counter: ", i)

// 		if i == 10 {
// 			break
// 		}

// 	}
// 	cancel() // membatalkan context

// 	time.Sleep(4 * time.Second)

// 	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
// }

// using wait group
func TestContextWithCancels(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	// Use WaitGroup to wait for goroutines to complete.
	wg := sync.WaitGroup{}
	destination := CreateCounter(ctx, &wg)

	wg.Add(1) // Add a WaitGroup counter for the goroutine reading the channel.
	go func() {
		defer wg.Done() // Signal that the goroutine is done.

		for i := range destination {
			fmt.Println("Counter: ", i)
			if i == 10 {
				cancel() // Cancel the context.
				break
			}
		}
	}()

	wg.Wait() // Wait for the goroutine to finish.

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 4*time.Second)

	defer cancel()

	// Use WaitGroup to wait for goroutines to complete.
	wg := sync.WaitGroup{}
	destination := CreateCounter(ctx, &wg)

	wg.Add(1) // Add a WaitGroup counter for the goroutine reading the channel.
	go func() {
		defer wg.Done() // Signal that the goroutine is done.

		for i := range destination {
			fmt.Println("Counter: ", i)
		}
	}()

	wg.Wait() // Wait for the goroutine to finish.

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}

func TestContextWithDeeadline(t *testing.T) {
	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))

	defer cancel()

	// Use WaitGroup to wait for goroutines to complete.
	wg := sync.WaitGroup{}
	destination := CreateCounter(ctx, &wg)

	wg.Add(1) // Add a WaitGroup counter for the goroutine reading the channel.
	go func() {
		defer wg.Done() // Signal that the goroutine is done.

		for i := range destination {
			fmt.Println("Counter: ", i)
		}
	}()

	wg.Wait() // Wait for the goroutine to finish.

	fmt.Println("Total Goroutine: ", runtime.NumGoroutine())
}
