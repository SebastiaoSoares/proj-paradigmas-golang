#define _POSIX_C_SOURCE 200809L

#include <pthread.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

enum {
    WORKER_COUNT = 4,
    JOB_COUNT = 8,
    QUEUE_CAPACITY = 3,
};

typedef struct {
    int id;
    int start;
    int end;
} Job;

typedef struct {
    Job job;
    int count;
} Result;

typedef struct {
    Job items[QUEUE_CAPACITY];
    size_t head;
    size_t tail;
    size_t count;
    bool closed;
    pthread_mutex_t mutex;
    pthread_cond_t not_empty;
    pthread_cond_t not_full;
} JobQueue;

typedef struct {
    JobQueue *queue;
    Result *results;
    pthread_mutex_t *output_mutex;
    size_t worker_id;
    bool visual;
} WorkerContext;

static const Job jobs[JOB_COUNT] = {
    {.id = 0, .start = 2, .end = 25000},
    {.id = 1, .start = 25001, .end = 50000},
    {.id = 2, .start = 50001, .end = 75000},
    {.id = 3, .start = 75001, .end = 100000},
    {.id = 4, .start = 100001, .end = 125000},
    {.id = 5, .start = 125001, .end = 150000},
    {.id = 6, .start = 150001, .end = 175000},
    {.id = 7, .start = 175001, .end = 200000},
};

static void check_pthread(int code, const char *operation) {
    if (code != 0) {
        fprintf(stderr, "%s: %s\n", operation, strerror(code));
        exit(EXIT_FAILURE);
    }
}

static bool is_prime(int value) {
    if (value < 2) {
        return false;
    }
    for (int divisor = 2; divisor <= value / divisor; divisor++) {
        if (value % divisor == 0) {
            return false;
        }
    }
    return true;
}

static int count_primes(int start, int end) {
    int count = 0;
    for (int candidate = start; candidate <= end; candidate++) {
        if (is_prime(candidate)) {
            count++;
        }
    }
    return count;
}

static void queue_push(JobQueue *queue, Job job) {
    check_pthread(pthread_mutex_lock(&queue->mutex), "pthread_mutex_lock");

    while (queue->count == QUEUE_CAPACITY) {
        check_pthread(
            pthread_cond_wait(&queue->not_full, &queue->mutex),
            "pthread_cond_wait(not_full)"
        );
    }

    queue->items[queue->tail] = job;
    queue->tail = (queue->tail + 1) % QUEUE_CAPACITY;
    queue->count++;

    check_pthread(pthread_cond_signal(&queue->not_empty), "pthread_cond_signal(not_empty)");
    check_pthread(pthread_mutex_unlock(&queue->mutex), "pthread_mutex_unlock");
}

static bool queue_pop(JobQueue *queue, Job *job) {
    check_pthread(pthread_mutex_lock(&queue->mutex), "pthread_mutex_lock");

    while (queue->count == 0 && !queue->closed) {
        check_pthread(
            pthread_cond_wait(&queue->not_empty, &queue->mutex),
            "pthread_cond_wait(not_empty)"
        );
    }

    if (queue->count == 0 && queue->closed) {
        check_pthread(pthread_mutex_unlock(&queue->mutex), "pthread_mutex_unlock");
        return false;
    }

    *job = queue->items[queue->head];
    queue->head = (queue->head + 1) % QUEUE_CAPACITY;
    queue->count--;

    check_pthread(pthread_cond_signal(&queue->not_full), "pthread_cond_signal(not_full)");
    check_pthread(pthread_mutex_unlock(&queue->mutex), "pthread_mutex_unlock");
    return true;
}

static void queue_close(JobQueue *queue) {
    check_pthread(pthread_mutex_lock(&queue->mutex), "pthread_mutex_lock");
    queue->closed = true;
    check_pthread(pthread_cond_broadcast(&queue->not_empty), "pthread_cond_broadcast");
    check_pthread(pthread_mutex_unlock(&queue->mutex), "pthread_mutex_unlock");
}

static void visual_event(
    WorkerContext *context,
    const char *phase,
    Job job,
    int count
) {
    if (!context->visual) {
        return;
    }

    check_pthread(
        pthread_mutex_lock(context->output_mutex),
        "pthread_mutex_lock(output)"
    );
    if (strcmp(phase, "start") == 0) {
        printf(
            "  P%zu  ▶ recebeu faixa %d [%d, %d]\n",
            context->worker_id,
            job.id + 1,
            job.start,
            job.end
        );
    } else {
        printf(
            "  P%zu  ✓ concluiu faixa %d: %d primos\n",
            context->worker_id,
            job.id + 1,
            count
        );
    }
    fflush(stdout);
    check_pthread(
        pthread_mutex_unlock(context->output_mutex),
        "pthread_mutex_unlock(output)"
    );
}

static void *worker(void *argument) {
    WorkerContext *context = argument;
    Job job;

    while (queue_pop(context->queue, &job)) {
        visual_event(context, "start", job, 0);
        if (context->visual) {
            const struct timespec pause = {.tv_sec = 0, .tv_nsec = 250000000L};
            (void)nanosleep(&pause, NULL);
        }

        Result result = {
            .job = job,
            .count = count_primes(job.start, job.end),
        };
        visual_event(context, "done", job, result.count);
        context->results[job.id] = result;
    }

    return NULL;
}

int main(int argc, char *argv[]) {
    bool visual = argc == 2 && strcmp(argv[1], "--visual") == 0;
    if (argc > 1 && !visual) {
        fprintf(stderr, "Uso: %s [--visual]\n", argv[0]);
        return EXIT_FAILURE;
    }

    JobQueue queue = {
        .mutex = PTHREAD_MUTEX_INITIALIZER,
        .not_empty = PTHREAD_COND_INITIALIZER,
        .not_full = PTHREAD_COND_INITIALIZER,
    };
    Result results[JOB_COUNT] = {0};
    pthread_t threads[WORKER_COUNT];
    WorkerContext contexts[WORKER_COUNT];
    pthread_mutex_t output_mutex = PTHREAD_MUTEX_INITIALIZER;

    if (visual) {
        printf("\nC · pthreads + mutex + variáveis de condição\n");
        printf("fila circular ──▶ 4 pthreads ──▶ vetor de resultados\n");
        printf("A pausa é didática; a contagem e a distribuição são reais.\n\n");
    }

    for (size_t index = 0; index < WORKER_COUNT; index++) {
        contexts[index] = (WorkerContext){
            .queue = &queue,
            .results = results,
            .output_mutex = &output_mutex,
            .worker_id = index + 1,
            .visual = visual,
        };
        check_pthread(
            pthread_create(&threads[index], NULL, worker, &contexts[index]),
            "pthread_create"
        );
    }

    for (size_t index = 0; index < JOB_COUNT; index++) {
        queue_push(&queue, jobs[index]);
    }
    queue_close(&queue);

    for (size_t index = 0; index < WORKER_COUNT; index++) {
        check_pthread(pthread_join(threads[index], NULL), "pthread_join");
    }

    int total = 0;
    if (visual) {
        printf("\n");
    }
    printf(
        "Comparativo: contagem de primos em %d faixas com %d workers\n",
        JOB_COUNT,
        WORKER_COUNT
    );
    for (size_t index = 0; index < JOB_COUNT; index++) {
        printf(
            "Faixa %d [%d, %d]: %d primos\n",
            results[index].job.id + 1,
            results[index].job.start,
            results[index].job.end,
            results[index].count
        );
        total += results[index].count;
    }
    printf("Total: %d primos entre 2 e 200000\n", total);

    check_pthread(pthread_cond_destroy(&queue.not_full), "pthread_cond_destroy");
    check_pthread(pthread_cond_destroy(&queue.not_empty), "pthread_cond_destroy");
    check_pthread(pthread_mutex_destroy(&queue.mutex), "pthread_mutex_destroy");
    check_pthread(pthread_mutex_destroy(&output_mutex), "pthread_mutex_destroy(output)");
    return EXIT_SUCCESS;
}
