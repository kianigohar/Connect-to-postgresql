package main

// https://data-nerd.blog/category/go/go-lang/
// modifying data and using transactions
import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "r"
	dbname   = "dvdrental"
)

var (
	ename string
	sal   int
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// First You begin a transaction with a call to db.Begin()
	// `tx` is an instance of `*sql.Tx` through which we can execute our queries
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Second, execute SQL queries against on the transaction instance. Note these are not applied to the database yet
	_, err = tx.ExecContext(ctx, "INSERT INTO public.emp (ename, email) VALUES('Smith', 'smith@acme.com')")
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
		return
	}

	// The next query is handled similarly
	_, err = tx.ExecContext(ctx, "UPDATE emp SET sal = 0 WHERE ename = 'Smith'")
	if err != nil {
		tx.Rollback()
		fmt.Println("\n", (err), "\n ....Transaction rollback!\n")
		return
	}

	// close the transaction with a Commit() or Rollback() method on the resulting Tx variable.
	// this applies the above changes to our database
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("....Transaction committed\n")
	}
}
