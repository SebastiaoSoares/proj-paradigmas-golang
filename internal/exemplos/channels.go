package exemplos

import (
	"fmt"
	"time"
)

// DemonstrarChannels demonstra o padrão Producer-Consumer com cinco pedidos.
func DemonstrarChannels() {
	// O channel transporta com segurança os IDs inteiros entre produtor e
	// consumidor. Como ele não é bufferizado, cada envio aguarda um recebimento.
	pedidos := make(chan int)

	// O produtor envia exatamente cinco pedidos e fecha o channel ao terminar.
	go func() {
		for pedido := 1; pedido <= 5; pedido++ {
			fmt.Printf("Produtor preparando o pedido %d...\n", pedido)
			// A pausa simula o tempo necessário para preparar cada pedido.
			time.Sleep(800 * time.Millisecond)

			// A mensagem identifica o momento em que o produtor disponibiliza o
			// pedido para o consumidor.
			fmt.Printf("Pedido enviado: %d\n", pedido)
			pedidos <- pedido
		}
		close(pedidos)
	}()

	// O consumidor usa range para receber todos os pedidos até o channel ser
	// fechado, evitando uma leitura indefinida ou uma tentativa insegura.
	for pedido := range pedidos {
		fmt.Printf("Pedido recebido: %d\n", pedido)
		// A pausa simula o consumidor processando o pedido antes de receber
		// o próximo valor do channel.
		time.Sleep(500 * time.Millisecond)
	}
}
