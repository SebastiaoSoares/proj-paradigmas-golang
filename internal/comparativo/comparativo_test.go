package comparativo

import (
	"os"
	"testing"
)

func TestCaminhoScriptApontaParaArquivoExistente(t *testing.T) {
	caminho, err := caminhoScript()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(caminho); err != nil {
		t.Fatalf("script não encontrado em %q: %v", caminho, err)
	}
}
