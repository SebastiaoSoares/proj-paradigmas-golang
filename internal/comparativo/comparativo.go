// Package comparativo integra o comparativo reproduzível ao menu principal.
package comparativo

import (
	"fmt"
	"os"
	"os/exec"
	"paradigmas_golang/internal/terminal"
	"path/filepath"
	"runtime"
)

// Executar apresenta o fluxo visual e depois testa e compara as saídas determinísticas.
func Executar() {
	terminal.Titulo("COMPARATIVO · GO, C E PYTHON", "Visualização dos workers seguida da verificação integral")

	if err := executarScript("demonstrar.sh", "--auto"); err != nil {
		terminal.Erro("Demonstração visual falhou: %v", err)
		return
	}
	if err := executarScript("verificar.sh"); err != nil {
		terminal.Erro("Verificação do comparativo falhou: %v", err)
		return
	}
	terminal.Sucesso("Comparativo completo aprovado")
}

func diretorioScripts() (string, error) {
	_, arquivoAtual, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("diretório do pacote indisponível")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(arquivoAtual), "..", "..", "examples", "comparativo")), nil
}

func executarScript(nome string, argumentos ...string) error {
	diretorio, err := diretorioScripts()
	if err != nil {
		return err
	}
	comando := exec.Command("sh", append([]string{filepath.Join(diretorio, nome)}, argumentos...)...)
	comando.Stdin = os.Stdin
	comando.Stdout = os.Stdout
	comando.Stderr = os.Stderr
	return comando.Run()
}
