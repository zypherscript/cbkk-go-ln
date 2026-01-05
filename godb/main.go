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

	//find all
	ppls, err := GetPeoples()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ppl := range ppls {
		fmt.Println(ppl)
	}

	//query by id
	ppl, err := GetPeople(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(*ppl)
}

func GetPeoples() ([]People, error) {
	// if err := db.Ping(); err != nil { //short if err but err only in if scope
	// 	return nil, err
	// }
	err := db.Ping()
	if err != nil {
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

func GetPeople(id int) (*People, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, name from people where id = $1" //mysql ? //sqlserver @id
	var ppl People
	//long
	// row := db.QueryRow(query, id)
	// err = row.Scan(&ppl.Id, &ppl.Name)
	// if err != nil {
	// 	return nil, err
	// }
	err = db.QueryRow(query, id).Scan(&ppl.Id, &ppl.Name)
	if err != nil {
		return nil, err
	}

	return &ppl, nil
}
