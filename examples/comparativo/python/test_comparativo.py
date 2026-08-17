import sys
import unittest
from pathlib import Path


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


if __name__ == "__main__":
    unittest.main()
