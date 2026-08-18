import sys
import unittest
from pathlib import Path
from typing import Optional


sys.path.insert(0, str(Path(__file__).resolve().parent))

from comparativo import Job, execute, is_prime  # noqa: E402


class IsPrimeTest(unittest.TestCase):
    def test_classifies_prime_and_non_prime_values(self) -> None:
        cases = (
            (-1, False),
            (0, False),
            (1, False),
            (2, True),
            (25, False),
            (97, True),
        )

        for value, expected in cases:
            with self.subTest(value=value):
                self.assertEqual(is_prime(value), expected)


class ExecuteTest(unittest.TestCase):
    def test_processes_jobs_and_orders_results_by_id(self) -> None:
        jobs = (
            Job(id=0, start=2, end=10),
            Job(id=1, start=11, end=20),
        )

        results = execute(2, jobs)

        self.assertEqual([result.count for result in results], [4, 4])
        self.assertEqual([result.job.id for result in results], [0, 1])

    def test_rejects_non_positive_worker_count(self) -> None:
        with self.assertRaisesRegex(ValueError, "maior que zero"):
            execute(0, ())

    def test_reports_worker_lifecycle_to_observer(self) -> None:
        jobs = (
            Job(id=0, start=2, end=10),
            Job(id=1, start=11, end=20),
        )
        events: list[tuple[int, str, int, Optional[int]]] = []

        execute(
            2,
            jobs,
            observer=lambda worker_id, phase, job, count: events.append(
                (worker_id, phase, job.id, count)
            ),
        )

        self.assertEqual(
            sorted((phase, job_id) for _, phase, job_id, _ in events),
            [("done", 0), ("done", 1), ("start", 0), ("start", 1)],
        )
        self.assertTrue(all(worker_id > 0 for worker_id, _, _, _ in events))


if __name__ == "__main__":
    unittest.main()
