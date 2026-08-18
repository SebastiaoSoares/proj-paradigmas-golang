#!/bin/sh

set -eu

repo_dir=$(CDPATH= cd -- "$(dirname -- "$0")/../.." && pwd)
build_dir=$(mktemp -d)
trap 'rm -rf "$build_dir"' EXIT HUP INT TERM

cd "$repo_dir"

go test -race ./examples/comparativo/go
go run ./examples/comparativo/go >"$build_dir/go.out"

python3 -B -m unittest discover \
    -s examples/comparativo/python \
    -p 'test_*.py'
python3 -B examples/comparativo/python/comparativo.py >"$build_dir/python.out"

cc \
    -std=c11 \
    -Wall \
    -Wextra \
    -Wpedantic \
    -Werror \
    -O2 \
    -pthread \
    examples/comparativo/c/main.c \
    -o "$build_dir/comparativo-c"

"$build_dir/comparativo-c" >"$build_dir/c.out"

printf '\n=== SAÍDA GO (goroutines e channels) ===\n'
cat "$build_dir/go.out"
printf '\n=== SAÍDA C (pthreads) ===\n'
cat "$build_dir/c.out"
printf '\n=== SAÍDA PYTHON (threads e Queue) ===\n'
cat "$build_dir/python.out"

printf '\n=== COMPARAÇÃO DAS SAÍDAS ===\n'
diff -u "$build_dir/go.out" "$build_dir/c.out"
diff -u "$build_dir/go.out" "$build_dir/python.out"

printf '%s\n' \
    "Go x C: saídas iguais." \
    "Go x Python: saídas iguais." \
    "Comparativo verificado: testes passaram e as três implementações produziram o mesmo resultado."
