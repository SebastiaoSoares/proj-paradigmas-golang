// Package exemplos reúne demonstrações pequenas de recursos de concorrência
// da linguagem Go.
package exemplos

import (
	"fmt"
	"sync"
	"time"
)

// DemonstrarWaitGroups simula três cozinheiros trabalhando em paralelo.
func DemonstrarWaitGroups() {
	// O WaitGroup acompanha quantas goroutines ainda precisam terminar.
	var wg sync.WaitGroup

	pratos := []string{"burger", "pizza", "salada"}

	for i, prato := range pratos {
		// Cada goroutine adicionada precisa ser registrada antes de começar.
		wg.Add(1)
		fmt.Printf("Atribuindo cozinheiro %d ao pedido de %s\n", i+1, prato)
		// A pausa torna cada atribuição visível no terminal.
		time.Sleep(400 * time.Millisecond)

		// Os parâmetros são passados para evitar que a goroutine dependa das
		// variáveis do loop depois que a próxima iteração começar.
		go func(cozinheiro int, prato string) {
			// Done sinaliza ao WaitGroup que este cozinheiro terminou.
			defer wg.Done()

			fmt.Printf("Cozinheiro %d começou a preparar %s\n", cozinheiro, prato)
			// A pausa torna visível a execução concorrente dos três cozinheiros.
			time.Sleep(500 * time.Millisecond)
			fmt.Printf("Cozinheiro %d finalizou %s\n", cozinheiro, prato)
		}(i+1, prato)
	}

	fmt.Println("Os três cozinheiros estão trabalhando. WaitGroup aguardando...")

	// Wait bloqueia a função até as três goroutines executarem Done.
	wg.Wait()
	fmt.Println("Todos os cozinheiros finalizaram. WaitGroup liberado.")
}
