package autenticacao

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken retorna um token assinado com as permissões do usuário
func CriarToken(usuarioID uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID
	// secret para assinar o token e garantir a autenticidade dele
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes) // método de assinatura = ES256
	return token.SignedString([]byte(config.SecretKey))            // secret
}
