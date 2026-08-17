from dataclasses import dataclass
from threading import Thread
from typing import Optional
from queue import Queue


WORKER_COUNT = 4


@dataclass(frozen=True)
class Job:
    id: int
    start: int
    end: int


@dataclass(frozen=True)
class Result:
    job: Job
    count: int


JOBS = (
    Job(id=0, start=2, end=25_000),
    Job(id=1, start=25_001, end=50_000),
    Job(id=2, start=50_001, end=75_000),
    Job(id=3, start=75_001, end=100_000),
    Job(id=4, start=100_001, end=125_000),
    Job(id=5, start=125_001, end=150_000),
    Job(id=6, start=150_001, end=175_000),
    Job(id=7, start=175_001, end=200_000),
)


def main() -> None:
    results = execute(WORKER_COUNT, JOBS)
    total = 0

    print(
        f"Comparativo: contagem de primos em {len(JOBS)} "
        f"faixas com {WORKER_COUNT} workers"
    )
    for result in results:
        print(
            f"Faixa {result.job.id + 1} "
            f"[{result.job.start}, {result.job.end}]: {result.count} primos"
        )
        total += result.count
    print(f"Total: {total} primos entre 2 e 200000")


def execute(workers: int, input_jobs: tuple[Job, ...]) -> list[Result]:
    if workers <= 0:
        raise ValueError("workers deve ser maior que zero")

    jobs_queue: Queue[Optional[Job]] = Queue(maxsize=workers)
    results_queue: Queue[Result] = Queue()
    threads = [
        Thread(
            target=worker,
            args=(jobs_queue, results_queue),
            name=f"worker-{index + 1}",
        )
        for index in range(workers)
    ]

    for thread in threads:
        thread.start()

    for current_job in input_jobs:
        jobs_queue.put(current_job)
    for _ in threads:
        jobs_queue.put(None)

    results = [results_queue.get() for _ in input_jobs]
    jobs_queue.join()

    for thread in threads:
        thread.join()

    return sorted(results, key=lambda result: result.job.id)


def worker(
    jobs_queue: Queue[Optional[Job]],
    results_queue: Queue[Result],
) -> None:
    while True:
        current_job = jobs_queue.get()
        try:
            if current_job is None:
                return
            results_queue.put(
                Result(
                    job=current_job,
                    count=count_primes(current_job.start, current_job.end),
                )
            )
        finally:
            jobs_queue.task_done()


def count_primes(start: int, end: int) -> int:
    return sum(1 for candidate in range(start, end + 1) if is_prime(candidate))


def is_prime(value: int) -> bool:
    if value < 2:
        return False
    divisor = 2
    while divisor <= value // divisor:
        if value % divisor == 0:
            return False
        divisor += 1
    return True


if __name__ == "__main__":
    main()
