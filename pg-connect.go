package main

// connecting to a PostgreSQL database with Go's database/sql package
import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {

	/*
	   variables required for connection string: connStr

	   user= (using default user for postgres database)
	   dbname= (using default database that comes with postgres)
	   password = (password used during initial setup)
	   host = (hostname or IP Address of server)
	   sslmode = (must be set to disabled unless using SSL)
	*/

	// https://data-nerd.blog/2020/04/11/connecting-to-postgresql-from-go-lang-project/
	//har 2 connection string check shod va ok ast
	// connStr := "user=postgres dbname=dvdrental password=r host=localhost sslmode=disable"
	connStr := "user=akiani dbname=dvdrental password=r host=localhost sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("\nSuccessfully connected to database!\n")
	fmt.Println("-0--------")
	////////////////////////////////////////////////////////////////////////
	var name string
	err = db.QueryRow("SELECT  address FROM address ORDER BY address_id ASC  LIMIT 1;").Scan(&name)
	if err != nil {
		panic(err)
	}
	fmt.Println(name)
	fmt.Println("-1--------")

	/////////////////////////////////////////////////////////////
	//https://data-nerd.blog/2020/04/25/querying-rows-from-postgresql-from-go-lang-project/
	// query rows from a table
	var (
		address  string
		district string
	)

	rows, err := db.Query("SELECT  address, district FROM address ORDER BY address_id ASC  LIMIT 3;")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&district, &address)
		if err != nil {
			panic(err)
		}
		fmt.Println("\n", district, " | ", address)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("-2--------")

	/////////////////////////////////////////////////////////////

}
