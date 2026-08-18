// Package interop demonstra a interoperabilidade entre Go e Python.
package interop

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// ExecutarPython envia texto para um script Python e imprime a saída recebida.
func ExecutarPython(texto string) error {
	_, arquivoAtual, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("não foi possível localizar o pacote de interoperabilidade")
	}

	caminhoScript := filepath.Join(filepath.Dir(arquivoAtual), "script.py")
	cmd := exec.Command("python3", caminhoScript, texto)

	saida, err := cmd.Output()
	if err != nil {
		if erroExecucao, ok := err.(*exec.ExitError); ok {
			mensagem := strings.TrimSpace(string(erroExecucao.Stderr))
			if mensagem != "" {
				return fmt.Errorf("script Python falhou: %s: %w", mensagem, err)
			}
		}
		return fmt.Errorf("não foi possível executar o script Python: %w", err)
	}

	fmt.Println(strings.TrimSuffix(string(saida), "\n"))
	return nil
}
