import sys


def main() -> None:
    if len(sys.argv) < 2:
        print("Erro: texto não fornecido.", file=sys.stderr)
        sys.exit(1)
    texto = sys.argv[1]
    print(texto.upper())


if __name__ == "__main__":
    main()
