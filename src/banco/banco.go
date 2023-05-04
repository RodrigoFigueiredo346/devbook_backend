package banco

import (
	"api_curso/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Driver
)

// Conectar abre a conexão com o banco de dados e a retorna
func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.StringConexaoBanco)
	if erro != nil {
		fmt.Println("nao abriu conexao...")
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		fmt.Println("nao pingou")
		db.Close()
		return nil, erro
	}

	return db, nil

}
