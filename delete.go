package main

// inserting records into a PostgreSQL database with Go's database/sql package
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

	////////////////////////
	// delete row
	sqlStatementDel := `
    DELETE FROM emp WHERE ename like $1;`
	res1, err := db.Exec(sqlStatementDel, "%")
	if err != nil {
		panic(err)
	}
	count, err := res1.RowsAffected()

	if err != nil {
		panic(err)
	}
	fmt.Printf("rows deleted: %v\n", count)
}
