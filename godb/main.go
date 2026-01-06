package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type People struct {
	Id   int
	Name string
}

var conn *pgx.Conn
var ctx = context.Background()

func main() {
	connStr := "postgres://postgres:postgres@localhost:5432/testdb?sslmode=disable"
	var err error
	conn, err = pgx.Connect(ctx, connStr)
	if err != nil {
		panic(err)
	}
	defer conn.Close(ctx)

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
	var n int
	err := conn.QueryRow(ctx, "SELECT 1").Scan(&n)
	if err != nil {
		return nil, err
	}

	query := "select id, name from people"
	rows, err := conn.Query(ctx, query)
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
	// err := db.Ping()
	// if err != nil {
	// 	return nil, err
	// }
	tx, err := conn.Begin(ctx)
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
	err = tx.QueryRow(ctx, query, id).Scan(&ppl.Id, &ppl.Name)
	if err != nil {
		tx.Rollback(ctx)
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &ppl, nil
}

func AddPeople(ppl People) error {
	// err := db.Ping()
	// if err != nil {
	// 	return err
	// }

	query := "insert into people (id, name) values ($1, $2)"
	result, err := conn.Exec(ctx, query, ppl.Id, ppl.Name)
	if err != nil {
		return err
	}
	affected := result.RowsAffected()
	if affected <= 0 {
		return errors.New("cannot insert")
	}

	return nil
}

func UpdatePeople(ppl People) error {
	// err := db.Ping()
	// if err != nil {
	// 	return err
	// }

	query := "update people set name=$2 where id=$1"
	result, err := conn.Exec(ctx, query, ppl.Id, ppl.Name)
	if err != nil {
		return err
	}
	affected := result.RowsAffected()
	if affected <= 0 {
		return errors.New("cannot update")
	}

	return nil
}

func DeletePeople(id int) error {
	// err := db.Ping()
	// if err != nil {
	// 	return err
	// }

	query := "delete from people where id=$1"
	result, err := conn.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	affected := result.RowsAffected()
	if affected <= 0 {
		return errors.New("cannot delete")
	}

	return nil
}
