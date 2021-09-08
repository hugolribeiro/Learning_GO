package estruturasmetodosinterfaces

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	retangulo := Retangulo{Largura: 10.0, Altura: 10.0}
	resultado := Perimetro(retangulo)
	esperado := 40.0

	if resultado != esperado {
		t.Errorf("resultado %.2f esperado %.2f", resultado, esperado)
	}
}

func TestArea(t *testing.T) {

	verificaArea := func(t *testing.T, forma Forma, esperado float64) {
		t.Helper()
		resultado := forma.Area()

		if resultado != esperado {
			t.Errorf("resultado %.2f, esperado %.2f", resultado, esperado)
		}
	}

	t.Run("retângulos", func(t *testing.T) {
		retangulo := Retangulo{Largura: 12.0, Altura: 6.0}
		esperado := 72.0
		verificaArea(t, retangulo, esperado)

	})

	t.Run("círculos", func(t *testing.T) {
		circulo := Circulo{10}
		esperado := 314.1592653589793
		verificaArea(t, circulo, esperado)

	})

}
