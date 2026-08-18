// Package terminal centraliza a identidade visual da aplicação no terminal.
package terminal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	reset   = "\033[0m"
	bold    = "\033[1m"
	dim     = "\033[2m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
)

func colorir(cor, texto string) string {
	if _, semCor := os.LookupEnv("NO_COLOR"); semCor {
		return texto
	}
	return cor + texto + reset
}

// Titulo imprime um cabeçalho visual comum a todos os módulos.
func Titulo(titulo, subtitulo string) {
	linha := strings.Repeat("═", 62)
	fmt.Printf("\n%s\n%s\n%s\n", colorir(cyan+bold, linha), colorir(cyan+bold, "  "+titulo), colorir(dim, "  "+subtitulo))
	fmt.Println(colorir(cyan+bold, linha))
}

func Info(formato string, args ...any)    { imprimir(blue, "●", formato, args...) }
func Passo(formato string, args ...any)   { imprimir(magenta, "→", formato, args...) }
func Sucesso(formato string, args ...any) { imprimir(green, "✓", formato, args...) }
func Alerta(formato string, args ...any)  { imprimir(yellow, "!", formato, args...) }
func Erro(formato string, args ...any)    { imprimir(red, "✕", formato, args...) }

func imprimir(cor, simbolo, formato string, args ...any) {
	mensagem := fmt.Sprintf(formato, args...)
	fmt.Printf("%s %s\n", colorir(cor+bold, simbolo), mensagem)
}

// Opcao destaca o número sem alterar o texto exigido pelo menu.
func Opcao(numero, descricao string) {
	fmt.Printf("  %s %s\n", colorir(cyan+bold, numero+"."), descricao)
}

func Prompt(texto string) { fmt.Print(colorir(yellow+bold, "❯ ") + texto) }

// Limpar redesenha o terminal usando sequências ANSI aceitas pelos terminais modernos.
func Limpar() {
	fmt.Print("\033[2J\033[H")
}

// Aguardar mantém o resultado visível até o usuário confirmar a próxima interação.
func Aguardar(leitor *bufio.Reader) {
	fmt.Println()
	Prompt("Pressione Enter para voltar ao menu...")
	_, _ = leitor.ReadString('\n')
}

// LerLinha lê inclusive espaços e remove somente os delimitadores da linha.
func LerLinha(leitor *bufio.Reader) (string, error) {
	linha, err := leitor.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	if err == io.EOF && len(linha) == 0 {
		return "", io.EOF
	}
	return strings.TrimRight(linha, "\r\n"), nil
}
