package iteracao

import "strings"

func Repetir(caractere string, quantidadeRepeticoes int) string {
	repeticoes := strings.Repeat(caractere, quantidadeRepeticoes)
	return repeticoes
}
