package arraysslices

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		resultado := Soma(numeros)
		esperado := 6

		if esperado != resultado {
			t.Errorf("resultado %d, esperado %d, dado %v", resultado, esperado, numeros)
		}
	})

}

func TestSomaTudo(t *testing.T) {

	resultado := SomaTudo([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}

	if !reflect.DeepEqual(resultado, esperado) {
		t.Errorf("resultado %v, esperado %v", resultado, esperado)
	}
}

func TestSomaTodoOResto(t *testing.T) {

	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper() // O Helper fará com que o número da linha relatará em caso de erro seja o da linha de fora dessa função, ou seja, a linha real de erro
		if !reflect.DeepEqual(resultado, esperado) {
			t.Errorf("resultado '%v', esperado '%v'", resultado, esperado)
		}
	}

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}

		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}

		verificaSomas(t, resultado, esperado)
	})
}
