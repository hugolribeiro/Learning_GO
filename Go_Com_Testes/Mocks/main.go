package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const inicioContagem = 3
const ultimaPalavra = "Vai!"

func Contagem(saida io.Writer) {
	for i := inicioContagem; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(saida, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(saida, ultimaPalavra)

}

func main() {
	Contagem(os.Stdout)
}
