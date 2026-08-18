package main

import (
	"fmt"
	"sync"
)

const workerCount = 4

type job struct {
	id    int
	start int
	end   int
}

type result struct {
	job   job
	count int
}

var jobs = []job{
	{id: 0, start: 2, end: 25_000},
	{id: 1, start: 25_001, end: 50_000},
	{id: 2, start: 50_001, end: 75_000},
	{id: 3, start: 75_001, end: 100_000},
	{id: 4, start: 100_001, end: 125_000},
	{id: 5, start: 125_001, end: 150_000},
	{id: 6, start: 150_001, end: 175_000},
	{id: 7, start: 175_001, end: 200_000},
}

func main() {
	results := execute(workerCount, jobs)
	total := 0

	fmt.Printf("Comparativo: contagem de primos em %d faixas com %d workers\n", len(jobs), workerCount)
	for _, result := range results {
		fmt.Printf("Faixa %d [%d, %d]: %d primos\n",
			result.job.id+1,
			result.job.start,
			result.job.end,
			result.count,
		)
		total += result.count
	}
	fmt.Printf("Total: %d primos entre 2 e 200000\n", total)
}

func execute(workers int, input []job) []result {
	jobsCh := make(chan job, workers)
	resultsCh := make(chan result, len(input))

	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go worker(jobsCh, resultsCh, &wg)
	}

	go func() {
		for _, currentJob := range input {
			jobsCh <- currentJob
		}
		close(jobsCh)
	}()

	go func() {
		wg.Wait()
		close(resultsCh)
	}()

	ordered := make([]result, len(input))
	for currentResult := range resultsCh {
		ordered[currentResult.job.id] = currentResult
	}

	return ordered
}

func worker(jobsCh <-chan job, resultsCh chan<- result, wg *sync.WaitGroup) {
	defer wg.Done()

	for currentJob := range jobsCh {
		resultsCh <- result{
			job:   currentJob,
			count: countPrimes(currentJob.start, currentJob.end),
		}
	}
}

func countPrimes(start, end int) int {
	count := 0
	for candidate := start; candidate <= end; candidate++ {
		if isPrime(candidate) {
			count++
		}
	}
	return count
}

func isPrime(value int) bool {
	if value < 2 {
		return false
	}
	for divisor := 2; divisor <= value/divisor; divisor++ {
		if value%divisor == 0 {
			return false
		}
	}
	return true
}
