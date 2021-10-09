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

	///////////////////////////////////////////////////////

	// insert a row
	sqlStatement := `INSERT INTO emp (ename, sal, email) 
    VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStatement, "Smith", 800, "smith@acme.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("\nRow inserted successfully!")
	}

	var (
		ename string
		email string
		sal   int
	)

	rows, err := db.Query("SELECT  ename,  email , sal FROM emp ;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&ename, &email, &sal)
		if err != nil {
			panic(err)
		}
		fmt.Println("\n", ename, " | ", email, " | ", sal)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("-2--------")
}
