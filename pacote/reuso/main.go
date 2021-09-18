package main

import (
	"fmt"

	"github.com/cod3rcursos/area"
)

func main() {
	fmt.Println(area.Circ(6.0))
	fmt.Println(area.Rect(5.0, 2.0))
	// fmt.Println(area._Triangulo(5.0, 2.0)) -> é privada!
}

// se tiver esse erro ao tentar utilizar módulos no $GOPATH
// ex: no required module provides package ... go.mod file not found in current directory or any parent directory; see 'go help modules'
// solução:
// $ go env -w GO111MODULE=off
// $ go env (para ver variáveis de ambiente do go)

// ------------------------------------------------------------
// Por exemplo, pode criar no github e baixar módulo:
// $ daniel@sim2701:~$ go get -u github.com/daniel-alec/goarea
