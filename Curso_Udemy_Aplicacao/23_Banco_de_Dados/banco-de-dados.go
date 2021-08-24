package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // o underline antes do pacote diz que em algum outro arquivo usaremos esse import
	"log"
)

func main() {
	// url de conexao
	// urlConexao := "usuario:senha@/banco"
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão está aberta ")

	// Para fazer queries usamos  db.Query
	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}
	defer linhas.Close()

	fmt.Println(linhas)

}
