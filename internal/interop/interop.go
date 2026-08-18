// Package interop demonstra a interoperabilidade entre Go e Python.
package interop

import (
	"fmt"
	"os"
	"os/exec"
	"paradigmas_golang/internal/terminal"
	"path/filepath"
	"runtime"
	"strings"
)

func localizarPython() (string, error) {
	candidatos := []string{"python3", "python"}
	if runtime.GOOS == "windows" {
		candidatos = []string{"python", "python3"}
	}

	for _, candidato := range candidatos {
		if caminho, err := exec.LookPath(candidato); err == nil {
			return caminho, nil
		}
	}
	return "", fmt.Errorf("Python 3 não encontrado; instale-o como 'python3' ou 'python'")
}

// ExecutarPython envia texto para um script Python e imprime a saída recebida.
func ExecutarPython(texto string) error {
	_, arquivoAtual, _, ok := runtime.Caller(0)
	if !ok {
		return fmt.Errorf("não foi possível localizar o pacote de interoperabilidade")
	}

	caminhoScript := filepath.Join(filepath.Dir(arquivoAtual), "script.py")
	python, err := localizarPython()
	if err != nil {
		return err
	}
	cmd := exec.Command(python, caminhoScript, texto)
	// Força uma codificação previsível também quando stdout é redirecionado no Windows.
	cmd.Env = append(os.Environ(), "PYTHONIOENCODING=utf-8")

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

	terminal.Sucesso("Python respondeu: %s", strings.TrimSpace(string(saida)))
	return nil
}
