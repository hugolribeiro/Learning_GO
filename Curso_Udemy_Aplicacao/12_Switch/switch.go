package main

import "fmt"

func diaDaSemana(numero int) string {
	// avalie a variável numero
	switch numero {
	// caso a variável numero equivale a 1 e assim por diante
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Domingo"
	default:
		return "Número inválido"
	}
}

//Outra maneira de utilizar o switch - mais utilizada quando queremos avaliar mais condições com mais variáveis também
func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	case numero == 3:
		diaDaSemana = "Terça-Feira"
	case numero == 4:
		diaDaSemana = "Quarta-Feira"
	case numero == 5:
		diaDaSemana = "Quinta-Feira"
	case numero == 6:
		diaDaSemana = "Sexta-Feira"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido"
	}
	return diaDaSemana
}

// uma palavra reservada dentro do Switch é o fallthrough, que faz pular algum caso e partir para o próximo

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(3)
	fmt.Println(dia)

	dia2 := diaDaSemana2(4)
	fmt.Println(dia2)
}
