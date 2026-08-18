package main

import (
	"bufio"
	"fmt"
	"os"
	"paradigmas_golang/internal/estudocaso"
	"paradigmas_golang/internal/exemplos"
	"paradigmas_golang/internal/interop"
	"strings"
)

func demonstrarInterop(leitor *bufio.Reader) {
	fmt.Println("\n=== Exemplo: Interoperabilidade Go + Python ===")
	fmt.Print("Digite um texto para converter em maiúsculas: ")

	texto, err := leitor.ReadString('\n')
	if err != nil {
		fmt.Printf("Erro ao ler o texto: %v\n", err)
		return
	}

	if err := interop.ExecutarPython(strings.TrimRight(texto, "\r\n")); err != nil {
		fmt.Printf("Erro na interoperabilidade com Python: %v\n", err)
	}
}

func main() {
	leitor := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nPROJETO DE PARADIGMAS DE PROGRAMAÇÃO: GO")
		fmt.Println("----------------------------------------")
		fmt.Println("1 - Simular cozinheiros (WaitGroup)")
		fmt.Println("2 - Simular pedidos (Channels)")
		fmt.Println("3 - Executar monitoramento de uptime")
		fmt.Println("4 - Executar todos os exemplos")
		fmt.Println("5 - Executar interoperabilidade Go + Python")
		fmt.Println("0 - Sair")
		fmt.Print("Escolha uma opção: ")

		opcao, err := leitor.ReadString('\n')
		if err != nil {
			return
		}

		switch strings.TrimSpace(opcao) {
		case "1":
			// Executa o exemplo de três workers sincronizados com WaitGroup.
			fmt.Println("\n=== Exemplo: WaitGroup ===")
			exemplos.DemonstrarWaitGroups()
		case "2":
			// Executa o exemplo de Producer-Consumer com um channel de inteiros.
			fmt.Println("\n=== Exemplo: Channels ===")
			exemplos.DemonstrarChannels()
		case "3":
			// Executa o estudo de caso de monitoramento de uptime.
			estudocaso.IniciarMonitoramento()
		case "4":
			// Executa os exemplos, o estudo de caso e a interoperabilidade.
			fmt.Println("\n=== Exemplo: WaitGroup ===")
			exemplos.DemonstrarWaitGroups()
			fmt.Println("\n=== Exemplo: Channels ===")
			exemplos.DemonstrarChannels()
			estudocaso.IniciarMonitoramento()
			demonstrarInterop(leitor)
		case "5":
			demonstrarInterop(leitor)
		case "0":
			fmt.Println("Programa encerrado.")
			return
		default:
			fmt.Println("Opção inválida. Escolha uma opção do menu.")
		}
	}
}
