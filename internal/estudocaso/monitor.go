package estudocaso

import (
	"fmt"
	"net/http"
	"paradigmas_golang/internal/terminal"
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

	terminal.Info("Disparando %d verificações concorrentes...", len(urls))

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()

			client := &http.Client{
				Timeout: 5 * time.Second,
			}

			inicio := time.Now()
			resp, err := client.Get(u)

			if err != nil {
				resultadosCh <- Resultado{URL: u, Erro: fmt.Errorf("após %s: %w", time.Since(inicio).Round(time.Millisecond), err)}
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
			terminal.Erro("%-36s indisponível (%v)", res.URL, res.Erro)
		} else {
			terminal.Sucesso("%-36s HTTP %d", res.URL, res.StatusCode)
		}
	}

	terminal.Sucesso("Monitoramento concluído; o channel de resultados foi fechado")
}
