package main

import "fmt"

var chanOwner = func() <-chan int {
	results := make(chan int, 5)
	go func() {
		defer close(results)
		for i := 1; i <= 5; i++ {
			results <- i
		}
	}()
	return results
}

var consumer = func(results <-chan int) {
	for result := range results {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}

func main() {
	chanOwner := chanOwner()
	consumer(chanOwner)
}
