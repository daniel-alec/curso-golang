package tabela

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parte: %s) - índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	t.Parallel()
	// para fazer mais testes utilizando um dataset
	testes := []struct {
		texto    string
		parte    string
		esperado int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"leonardo", "o", 2},
	}

	for _, teste := range testes {
		t.Logf("Massa: %v", teste) // log só será apresentando se algum registro não passar na validação
		atual := strings.Index(teste.texto, teste.parte)

		if atual != teste.esperado {
			t.Errorf(msgIndex, teste.texto, teste.parte, teste.esperado, atual)
		}
	}
}

// rodar teste verboso (linha de comando)
// cd testes/tabela
// go test -v
