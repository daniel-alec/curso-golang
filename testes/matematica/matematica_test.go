package matematica

import "testing"

const erroPadrao = "Valor esperado %v, mas o resultado encontrado foi %v."

func TestMedia(t *testing.T) { // Convenção Test...
	t.Parallel()
	valorEsperado := 7.28
	valor := Media(7.2, 9.9, 6.1, 5.9)

	if valor != valorEsperado {
		t.Errorf(erroPadrao, valorEsperado, valor)
	}
}

// Testar do terminal:
// cd testes/matematica
// go test
//
// ou
//
// go test ./testes/matematica/...

//--------------------------------------

// Testar cobertura do terminal:
// cd matematica/
// go test -cover

// Gera resultado:
// go test -coverprofile=resultado.out

// Detalhes sobre o resultado:
// go tool cover -func=resultado.out

// Gerar html do resulado
// go tool cover -html=resultado.out
