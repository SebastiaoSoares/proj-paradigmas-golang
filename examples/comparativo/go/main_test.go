package main

import "testing"

func TestIsPrime(t *testing.T) {
	tests := []struct {
		name  string
		value int
		want  bool
	}{
		{name: "negative", value: -1, want: false},
		{name: "zero", value: 0, want: false},
		{name: "one", value: 1, want: false},
		{name: "two", value: 2, want: true},
		{name: "composite", value: 25, want: false},
		{name: "prime", value: 97, want: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := isPrime(test.value); got != test.want {
				t.Fatalf("isPrime(%d) = %t; want %t", test.value, got, test.want)
			}
		})
	}
}

func TestExecute(t *testing.T) {
	input := []job{
		{id: 0, start: 2, end: 10},
		{id: 1, start: 11, end: 20},
	}

	got := execute(2, input)
	want := []int{4, 4}

	if len(got) != len(want) {
		t.Fatalf("execute returned %d results; want %d", len(got), len(want))
	}
	for index, expected := range want {
		if got[index].count != expected {
			t.Errorf("result %d = %d; want %d", index, got[index].count, expected)
		}
	}
}

func TestExecuteObservedReportsWorkerLifecycle(t *testing.T) {
	input := []job{
		{id: 0, start: 2, end: 10},
		{id: 1, start: 11, end: 20},
	}
	events := make(chan workerEvent, len(input)*2)

	executeObserved(2, input, func(event workerEvent) {
		events <- event
	}, 0)
	close(events)

	starts := 0
	done := 0
	for event := range events {
		switch event.phase {
		case "start":
			starts++
		case "done":
			done++
			if event.count != 4 {
				t.Errorf("job %d reported %d primes; want 4", event.job.id, event.count)
			}
		default:
			t.Errorf("unexpected event phase %q", event.phase)
		}
	}

	if starts != len(input) || done != len(input) {
		t.Fatalf("got %d start and %d done events; want %d of each", starts, done, len(input))
	}
}
