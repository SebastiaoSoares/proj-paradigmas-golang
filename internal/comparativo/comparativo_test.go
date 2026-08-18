package comparativo

import (
	"os"
	"testing"
)

func TestDiretorioContemScriptsDoComparativo(t *testing.T) {
	diretorio, err := diretorioScripts()
	if err != nil {
		t.Fatal(err)
	}
	for _, nome := range []string{"demonstrar.sh", "verificar.sh"} {
		caminho := diretorio + string(os.PathSeparator) + nome
		if _, err := os.Stat(caminho); err != nil {
			t.Fatalf("script não encontrado em %q: %v", caminho, err)
		}
	}
}
