package main

import (
	"bufio"
	"fmt"
	"os"
	"paradigmas_golang/internal/estudocaso"
	"paradigmas_golang/internal/exemplos"
	"strings"
)

func main() {
	leitor := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nPROJETO DE PARADIGMAS DE PROGRAMAÇÃO: GO")
		fmt.Println("----------------------------------------")
		fmt.Println("1 - Simular cozinheiros (WaitGroup)")
		fmt.Println("2 - Simular pedidos (Channels)")
		fmt.Println("3 - Executar monitoramento de uptime")
		fmt.Println("4 - Executar todos os exemplos")
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
			// Executa os dois exemplos e, em seguida, o estudo de caso.
			fmt.Println("\n=== Exemplo: WaitGroup ===")
			exemplos.DemonstrarWaitGroups()
			fmt.Println("\n=== Exemplo: Channels ===")
			exemplos.DemonstrarChannels()
			estudocaso.IniciarMonitoramento()
		case "0":
			fmt.Println("Programa encerrado.")
			return
		default:
			fmt.Println("Opção inválida. Escolha uma opção do menu.")
		}
	}
}
