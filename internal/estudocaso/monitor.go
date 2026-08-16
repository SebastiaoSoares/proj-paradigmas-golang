package estudocaso

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Resultado struct {
	URL        string
	StatusCode int
	Erro       error
}

func IniciarMonitoramento() {
	urls := []string{
		"https://go.dev",
		"https://github.com",
		"https://www.google.com",
		"https://httpbin.org/get",
		"https://www.cloudflare.com",
		"https://invalid-url-for-testing.local", 
	}

	resultadosCh := make(chan Resultado)
	
	var wg sync.WaitGroup

	fmt.Println("=== Iniciando Monitoramento de Uptime ===")

	for _, url := range urls {
		wg.Add(1) 
		
		go func(u string) {
			defer wg.Done() 

			client := &http.Client{
				Timeout: 5 * time.Second,
			}

			resp, err := client.Get(u)
			
			if err != nil {
				resultadosCh <- Resultado{URL: u, Erro: err}
				return
			}
			defer resp.Body.Close()

			resultadosCh <- Resultado{URL: u, StatusCode: resp.StatusCode, Erro: nil}
		}(url)
	}

	go func() {
		wg.Wait()          
		close(resultadosCh) 
	}()

	for res := range resultadosCh {
		if res.Erro != nil {
			fmt.Printf("[FALHA] URL: %s | Erro: %v\n", res.URL, res.Erro)
		} else {
			fmt.Printf("[ OK ] URL: %s | Status Code: %d\n", res.URL, res.StatusCode)
		}
	}

	fmt.Println("=== Monitoramento Concluído ===")
}