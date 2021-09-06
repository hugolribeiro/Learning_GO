package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	numeroRepeticoes := 5
	repeticoes := Repetir("a", numeroRepeticoes)
	esperado := "aaaaa"
	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}

func ExampleRepetir() {
	repeticao := Repetir("a", 6)
	fmt.Println(repeticao)
	// Output: aaaaaa
}
