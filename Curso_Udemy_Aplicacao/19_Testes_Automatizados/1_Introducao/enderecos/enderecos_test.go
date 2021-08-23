// TESTE DE UNIDADE

package enderecos

import "testing"

// Essa é a sintexe de uma função de teste
func TestTipoDeEnderecoAvenidaDeveriaRetornarAvenida(t *testing.T) {
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido)
	}
}
