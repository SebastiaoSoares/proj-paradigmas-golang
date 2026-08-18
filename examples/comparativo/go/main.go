package main

import (
	"fmt"
	"os"
	"sync"
	"time"
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

type workerEvent struct {
	workerID int
	phase    string
	job      job
	count    int
}

type observer func(workerEvent)

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
	visual := len(os.Args) == 2 && os.Args[1] == "--visual"
	if len(os.Args) > 1 && !visual {
		fmt.Fprintf(os.Stderr, "Uso: %s [--visual]\n", os.Args[0])
		os.Exit(2)
	}

	var observe observer
	var pause time.Duration
	if visual {
		fmt.Println("\nGO · goroutines + channels")
		fmt.Println("fila de jobs ──▶ 4 goroutines ──▶ channel de resultados")
		fmt.Println("A pausa é didática; a contagem e a distribuição são reais.")
		fmt.Println()

		var outputMutex sync.Mutex
		observe = func(event workerEvent) {
			outputMutex.Lock()
			defer outputMutex.Unlock()
			if event.phase == "start" {
				fmt.Printf("  G%d  ▶ recebeu faixa %d [%d, %d]\n",
					event.workerID, event.job.id+1, event.job.start, event.job.end)
				return
			}
			fmt.Printf("  G%d  ✓ concluiu faixa %d: %d primos\n",
				event.workerID, event.job.id+1, event.count)
		}
		pause = 250 * time.Millisecond
	}

	results := executeObserved(workerCount, jobs, observe, pause)
	total := 0

	if visual {
		fmt.Println()
	}
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
	return executeObserved(workers, input, nil, 0)
}

func executeObserved(workers int, input []job, observe observer, pause time.Duration) []result {
	jobsCh := make(chan job, workers)
	resultsCh := make(chan result, len(input))

	var wg sync.WaitGroup
	for workerIndex := range workers {
		wg.Add(1)
		go worker(workerIndex+1, jobsCh, resultsCh, &wg, observe, pause)
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

func worker(
	workerID int,
	jobsCh <-chan job,
	resultsCh chan<- result,
	wg *sync.WaitGroup,
	observe observer,
	pause time.Duration,
) {
	defer wg.Done()

	for currentJob := range jobsCh {
		if observe != nil {
			observe(workerEvent{workerID: workerID, phase: "start", job: currentJob})
		}
		if pause > 0 {
			time.Sleep(pause)
		}

		currentResult := result{
			job:   currentJob,
			count: countPrimes(currentJob.start, currentJob.end),
		}
		if observe != nil {
			observe(workerEvent{
				workerID: workerID,
				phase:    "done",
				job:      currentJob,
				count:    currentResult.count,
			})
		}
		resultsCh <- currentResult
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
