package main

import (
	"bufio"
	"io"
	"os"
	"paradigmas_golang/internal/terminal"
	"strings"
	"testing"
)

func capturarSaida(t *testing.T, executar func()) string {
	t.Helper()
	leitura, escrita, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	original := os.Stdout
	os.Stdout = escrita
	defer func() { os.Stdout = original }()
	executar()
	os.Stdout = original
	if err := escrita.Close(); err != nil {
		t.Fatal(err)
	}
	saida, err := io.ReadAll(leitura)
	if err != nil {
		t.Fatal(err)
	}
	if err := leitura.Close(); err != nil {
		t.Fatal(err)
	}
	return string(saida)
}

func TestMenuPercorreTodasAsOpcoes(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	var exemplos, interop, monitor, comparar int
	leitor := bufio.NewReader(strings.NewReader("invalida\n\n1\n\n2\ntexto com espaços\n\n3\n\n4\n\n0\n"))

	saida := capturarSaida(t, func() {
		executarMenu(leitor, acoesMenu{
			exemplos: func() { exemplos++ },
			interop: func(leitor *bufio.Reader) {
				interop++
				texto, err := terminal.LerLinha(leitor)
				if err != nil || texto != "texto com espaços" {
					t.Fatalf("entrada da interoperabilidade = %q, %v", texto, err)
				}
			},
			monitor:  func() { monitor++ },
			comparar: func() { comparar++ },
		})
	})

	if exemplos != 1 || interop != 1 || monitor != 1 || comparar != 1 {
		t.Fatalf("ações chamadas: exemplos=%d, interop=%d, monitor=%d, comparar=%d", exemplos, interop, monitor, comparar)
	}
	for _, trecho := range []string{"Opção \"invalida\" inválida", "ESTUDO DE CASO", "Programa encerrado"} {
		if !strings.Contains(saida, trecho) {
			t.Errorf("saída não contém %q", trecho)
		}
	}
}

func TestMenuEncerraAoReceberEOF(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	saida := capturarSaida(t, func() {
		executarMenu(bufio.NewReader(strings.NewReader("")), acoesMenu{
			exemplos: func() {}, interop: func(*bufio.Reader) {}, monitor: func() {}, comparar: func() {},
		})
	})
	if !strings.Contains(saida, "Não foi possível ler a opção") {
		t.Fatalf("mensagem de EOF ausente: %q", saida)
	}
}
