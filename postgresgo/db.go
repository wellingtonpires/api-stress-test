package postgresgo

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func openConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=apenasteste dbname=stress_test sslmode=disable")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	return db, err
}

func InsertData(apelido string, nome string, dataNasc string, stack []string) {
	con, err := openConnection()
	id := uuid.New()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer con.Close()
	sql := `INSERT INTO pessoas VALUES ($1, $2, $3, $4)`
	con.Exec(sql, id, apelido, nome, dataNasc)
}
