// Arquivo responsável pela conexão com o DB

package db

import (
	"api-go/configs"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	strConnect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disabled", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgress", strConnect) 
	if err != nil {
		// Não é bom usar panic em produção
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}