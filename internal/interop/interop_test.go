package interop

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestExecutarPythonImprimeTextoEmMaiusculas(t *testing.T) {
	diretorioOriginal, err := os.Getwd()
	if err != nil {
		t.Fatalf("não foi possível consultar o diretório atual: %v", err)
	}

	if err := os.Chdir(t.TempDir()); err != nil {
		t.Fatalf("não foi possível trocar o diretório atual: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(diretorioOriginal); err != nil {
			t.Errorf("não foi possível restaurar o diretório atual: %v", err)
		}
	})

	leitura, escrita, err := os.Pipe()
	if err != nil {
		t.Fatalf("não foi possível criar pipe para capturar stdout: %v", err)
	}

	stdoutOriginal := os.Stdout
	os.Stdout = escrita
	t.Cleanup(func() { os.Stdout = stdoutOriginal })

	errExecucao := ExecutarPython("Olá, Go e Python!")
	if err := escrita.Close(); err != nil {
		t.Fatalf("não foi possível fechar a escrita do pipe: %v", err)
	}
	os.Stdout = stdoutOriginal

	saida, err := io.ReadAll(leitura)
	if err != nil {
		t.Fatalf("não foi possível ler stdout: %v", err)
	}
	if err := leitura.Close(); err != nil {
		t.Fatalf("não foi possível fechar a leitura do pipe: %v", err)
	}

	if errExecucao != nil {
		t.Fatalf("ExecutarPython retornou erro: %v", errExecucao)
	}

	obtido := strings.TrimSpace(string(saida))
	esperado := "OLÁ, GO E PYTHON!"
	if obtido != esperado {
		t.Fatalf("saída inesperada: obtido %q, esperado %q", obtido, esperado)
	}
}
