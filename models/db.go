package models

import(
	"fmt"
	"database/sql"
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

func InsertData(s string) {
	fmt.Println("bateu")
	con, err := openConnection()
	if err != nil {
		fmt.Println("deu erro porra" + err.Error())
		return
	}
	defer con.Close()
	sql := `INSERT INTO pessoas VALUES ('uuid-2138', 'apelido', 'nomepessoa', '2022-03-02')`
	con.QueryRow(sql)
	fmt.Println("Chegou aqui carajo!!!")
}
