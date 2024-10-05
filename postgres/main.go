package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	_ "github.com/jackc/pgx/v4/stdlib"
)


func main() {
	dsn := url.URL{
		Scheme: "postgres",
		Host: "localhost:5432",
		User: url.UserPassword("postgres", "mysecretpassword"),
		Path: "gopost",
	}


	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	db , err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("Sql.OPen", err)
		return
	}

	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	row := db.QueryRowContext(context.Background(), "SELECT * FROM users WHERE name = 'amr'")
	if err := row.Err(); err != nil{
		fmt.Println("db query row context", err)
		return
	}
	var birth_year int64
	if err := row.Scan(&birth_year); err != nil {
		fmt.Println("row.Scan", err)
		return
	}
	fmt.Println("birth_year",birth_year)

}