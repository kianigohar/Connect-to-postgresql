package main

// updating PostgreSQL records using Go's database/sql package

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "r"
	dbname   = "dvdrental"
)

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// CREATE TABLE emp (
	// 	empno SERIAL PRIMARY KEY,
	// 	ename TEXT,
	// 	sal INT,
	// 	email TEXT UNIQUE NOT NULL
	//   );

	// update row
	sqlStatementUpdt := `
    UPDATE emp
    SET sal = 3000 WHERE ename = $1;`
	res, err := db.Exec(sqlStatementUpdt, "Smith")
	if err != nil {
		panic(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Printf("rows updated: %v\n", count)

	// delete row
	sqlStatementDel := `
    DELETE FROM emp
    WHERE ename = $1;`
	res1, err := db.Exec(sqlStatementDel, "Jones")
	if err != nil {
		panic(err)
	}
	count, err = res1.RowsAffected()

	if err != nil {
		panic(err)
	}
	fmt.Printf("rows deleted: %v\n", count)
}
