// Package comparativo integra o comparativo reproduzível ao menu principal.
package comparativo

import (
	"fmt"
	"os/exec"
	"paradigmas_golang/internal/terminal"
	"path/filepath"
	"runtime"
	"strings"
)

// Verificar executa o script que testa Go, C/pthreads e Python e compara suas saídas.
func Verificar() {
	terminal.Titulo("COMPARATIVO COMPLETO · GO, C E PYTHON", "Testes, execuções e comparação integral das saídas")

	caminho, err := caminhoScript()
	if err != nil {
		terminal.Erro("Não foi possível localizar o comparativo: %v", err)
		return
	}

	terminal.Info("Executando examples/comparativo/verificar.sh...")
	comando := exec.Command("sh", caminho)
	saida, err := comando.CombinedOutput()
	if len(saida) > 0 {
		fmt.Print(string(saida))
		if !strings.HasSuffix(string(saida), "\n") {
			fmt.Println()
		}
	}
	if err != nil {
		terminal.Erro("Comparativo reprovado: %v", err)
		return
	}
	terminal.Sucesso("Comparativo completo aprovado")
}

func caminhoScript() (string, error) {
	_, arquivoAtual, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("diretório do pacote indisponível")
	}
	caminho := filepath.Join(filepath.Dir(arquivoAtual), "..", "..", "examples", "comparativo", "verificar.sh")
	return filepath.Clean(caminho), nil
}
