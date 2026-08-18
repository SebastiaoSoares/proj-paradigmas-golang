#!/bin/sh

set -eu

if [ "$#" -gt 1 ] || { [ "$#" -eq 1 ] && [ "$1" != '--auto' ]; }; then
    printf 'Uso: %s [--auto]\n' "$0" >&2
    exit 2
fi

repo_dir=$(CDPATH= cd -- "$(dirname -- "$0")/../.." && pwd)
build_dir=$(mktemp -d)
trap 'rm -rf "$build_dir"' EXIT HUP INT TERM

if [ -t 1 ] && [ -z "${NO_COLOR:-}" ]; then
    bold='\033[1m'
    cyan='\033[36m'
    green='\033[32m'
    dim='\033[2m'
    reset='\033[0m'
else
    bold=''
    cyan=''
    green=''
    dim=''
    reset=''
fi

section() {
    printf '\n%b%s%b\n' "$cyan$bold" '══════════════════════════════════════════════════════════════' "$reset"
    printf '%b  %s%b\n' "$cyan$bold" "$1" "$reset"
    printf '%b%s%b\n' "$cyan$bold" '══════════════════════════════════════════════════════════════' "$reset"
}

continue_demo() {
    if [ "${1:-}" = '--auto' ] || [ ! -t 0 ]; then
        sleep 1
        return
    fi
    printf '\n%bPressione Enter para ver a próxima implementação...%b' "$dim" "$reset"
    read -r _
}

cd "$repo_dir"

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

section 'COMPARATIVO VISUAL · O MESMO WORKER POOL EM 3 LINGUAGENS'
printf '%s\n' \
    'Problema: contar primos em 8 faixas usando 4 workers.' \
    'Os identificadores mostram qual worker recebeu cada faixa.' \
    'A ordem pode mudar entre execuções: isso é concorrência observável.'

continue_demo "${1:-}"
section '1/3 · GO'
go run ./examples/comparativo/go --visual

continue_demo "${1:-}"
section '2/3 · C/POSIX'
"$build_dir/comparativo-c" --visual

continue_demo "${1:-}"
section '3/3 · PYTHON'
python3 -B examples/comparativo/python/comparativo.py --visual

section 'LEITURA FINAL'
printf '%-10s %-18s %-28s\n' 'Linguagem' 'Workers' 'Coordenação'
printf '%-10s %-18s %-28s\n' 'Go' 'goroutines' 'channels + WaitGroup'
printf '%-10s %-18s %-28s\n' 'C' 'pthreads' 'mutex + condições'
printf '%-10s %-18s %-28s\n' 'Python' 'Thread' 'Queue + join'
printf '\n%b✓ As três versões chegaram ao mesmo total: 17984 primos.%b\n' "$green$bold" "$reset"
printf '%bA pausa visual não representa desempenho e não deve ser usada como benchmark.%b\n' "$dim" "$reset"
