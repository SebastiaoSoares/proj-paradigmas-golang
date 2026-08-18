package main

import (
	"bufio"
	"fmt"
	"os"
	"paradigmas_golang/internal/estudocaso"
	"paradigmas_golang/internal/exemplos"
	"paradigmas_golang/internal/interop"
	"paradigmas_golang/internal/terminal"
)

func executarExemplosBasicos() {
	terminal.Titulo("EXEMPLO 1 · WAITGROUP", "Sincronização de goroutines trabalhando em paralelo")
	exemplos.DemonstrarWaitGroups()

	terminal.Titulo("EXEMPLO 2 · CHANNELS", "Comunicação segura no padrão produtor–consumidor")
	exemplos.DemonstrarChannels()
}

func executarInteroperabilidade(leitor *bufio.Reader) {
	terminal.Titulo("INTEROPERABILIDADE · GO + PYTHON", "Go inicia um processo Python e captura sua saída")
	terminal.Prompt("Digite um texto para converter em maiúsculas: ")
	texto, err := terminal.LerLinha(leitor)
	if err != nil {
		terminal.Erro("Não foi possível ler o texto.")
		return
	}
	if texto == "" {
		terminal.Alerta("Digite ao menos um caractere para realizar a conversão.")
		return
	}

	if err = interop.ExecutarPython(texto); err != nil {
		terminal.Erro("Interoperabilidade com Python: %v", err)
	}
}

func main() {
	leitor := bufio.NewReader(os.Stdin)

	for {
		terminal.Limpar()
		terminal.Titulo("GO · CONCORRÊNCIA NA PRÁTICA", "Goroutines, channels e interoperabilidade")
		terminal.Opcao("1", "Exemplos Basicos")
		terminal.Opcao("2", "Interoperabilidade Python")
		terminal.Opcao("3", "Estudo de Caso (Monitor de Uptime)")
		terminal.Opcao("0", "Sair")
		fmt.Println()
		terminal.Prompt("Escolha uma opção: ")

		opcao, err := terminal.LerLinha(leitor)
		if err != nil {
			terminal.Erro("Não foi possível ler a opção.")
			return
		}

		switch opcao {
		case "1":
			terminal.Limpar()
			executarExemplosBasicos()
			terminal.Aguardar(leitor)
		case "2":
			terminal.Limpar()
			executarInteroperabilidade(leitor)
			terminal.Aguardar(leitor)
		case "3":
			terminal.Limpar()
			terminal.Titulo("ESTUDO DE CASO · MONITOR DE UPTIME", "Requisições HTTP concorrentes com timeout")
			estudocaso.IniciarMonitoramento()
			terminal.Aguardar(leitor)
		case "0":
			terminal.Limpar()
			terminal.Sucesso("Programa encerrado. Até a próxima!")
			return
		default:
			terminal.Limpar()
			terminal.Alerta("Opção %q inválida. Escolha 1, 2, 3 ou 0.", opcao)
			terminal.Aguardar(leitor)
		}
	}
}
