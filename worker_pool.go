package main

import "fmt"

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- j * j
		fmt.Printf("Worker %d processed job %d\n", id, j)
	}
}

func showWorkerPool() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for r := 0; r < 9; r++ {
		fmt.Println("Result:", <-results)
	}
}
