package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type People struct {
	Id   int
	Name string
}

var db *sql.DB

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ppls, err := getPeoples()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ppl := range ppls {
		fmt.Println(ppl)
	}
}

func getPeoples() ([]People, error) {
	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := "select id, name from people"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ppls := []People{}
	for rows.Next() {
		ppl := People{}
		err = rows.Scan(&ppl.Id, &ppl.Name) //select id, name from ...
		if err != nil {
			return nil, err
		}
		ppls = append(ppls, ppl)
	}
	return ppls, nil
}
