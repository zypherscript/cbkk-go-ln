package main

import (
	"database/sql"
	"errors"
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

	//insert
	addPpl := People{3, "test"}
	err = AddPeople(addPpl)
	if err != nil {
		panic(err)
	} else {
		println("successfully added")
	}
	//find all
	ppls, err := GetPeoples()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ppl := range ppls {
		fmt.Println(ppl)
	}

	//update
	addPpl = People{3, "test2"}
	err = UpdatePeople(addPpl)
	if err != nil {
		panic(err)
	} else {
		println("successfully updated")
	}

	//find all
	ppls, err = GetPeoples()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ppl := range ppls {
		fmt.Println(ppl)
	}

	//query by id
	println(">>> try queryRow")
	ppl, err := GetPeople(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(*ppl)

	//update
	err = DeletePeople(addPpl.Id)
	if err != nil {
		panic(err)
	} else {
		println("successfully deleted")
	}
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

func AddPeople(ppl People) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	query := "insert into people (id, name) values ($1, $2)"
	result, err := db.Exec(query, ppl.Id, ppl.Name)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot insert")
	}

	return nil
}

func UpdatePeople(ppl People) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	query := "update people set name=$2 where id=$1"
	result, err := db.Exec(query, ppl.Id, ppl.Name)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot update")
	}

	return nil
}

func DeletePeople(id int) error {
	err := db.Ping()
	if err != nil {
		return err
	}

	query := "delete from people where id=$1"
	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot delete")
	}

	return nil
}
