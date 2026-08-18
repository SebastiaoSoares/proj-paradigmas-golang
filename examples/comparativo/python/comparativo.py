from dataclasses import dataclass
import sys
from threading import Lock, Thread
from time import sleep
from typing import Callable, Optional
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


Observer = Callable[[int, str, Job, Optional[int]], None]


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
    visual = sys.argv[1:] == ["--visual"]
    if len(sys.argv) > 1 and not visual:
        print(f"Uso: {sys.argv[0]} [--visual]", file=sys.stderr)
        raise SystemExit(2)

    observer: Optional[Observer] = None
    pause = 0.0
    if visual:
        print("\nPYTHON · Thread + Queue")
        print("fila de jobs ──▶ 4 threads ──▶ fila de resultados")
        print("A pausa é didática; a contagem e a distribuição são reais.\n")
        output_lock = Lock()

        def show_event(
            worker_id: int,
            phase: str,
            job: Job,
            count: Optional[int],
        ) -> None:
            with output_lock:
                show_event_line(worker_id, phase, job, count)

        def show_event_line(
            worker_id: int,
            phase: str,
            job: Job,
            count: Optional[int],
        ) -> None:
            if phase == "start":
                print(
                    f"  T{worker_id}  ▶ recebeu faixa {job.id + 1} "
                    f"[{job.start}, {job.end}]",
                    flush=True,
                )
                return
            print(
                f"  T{worker_id}  ✓ concluiu faixa {job.id + 1}: "
                f"{count} primos",
                flush=True,
            )

        observer = show_event
        pause = 0.25

    results = execute(WORKER_COUNT, JOBS, observer=observer, pause=pause)
    total = 0

    if visual:
        print()
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


def execute(
    workers: int,
    input_jobs: tuple[Job, ...],
    observer: Optional[Observer] = None,
    pause: float = 0.0,
) -> list[Result]:
    if workers <= 0:
        raise ValueError("workers deve ser maior que zero")

    jobs_queue: Queue[Optional[Job]] = Queue(maxsize=workers)
    results_queue: Queue[Result] = Queue()
    threads = [
        Thread(
            target=worker,
            args=(index + 1, jobs_queue, results_queue, observer, pause),
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
    worker_id: int,
    jobs_queue: Queue[Optional[Job]],
    results_queue: Queue[Result],
    observer: Optional[Observer],
    pause: float,
) -> None:
    while True:
        current_job = jobs_queue.get()
        try:
            if current_job is None:
                return
            if observer is not None:
                observer(worker_id, "start", current_job, None)
            if pause > 0:
                sleep(pause)

            result = Result(
                job=current_job,
                count=count_primes(current_job.start, current_job.end),
            )
            if observer is not None:
                observer(worker_id, "done", current_job, result.count)
            results_queue.put(result)
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
