package sql

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const (
	invalidIdError = "No receipt found for that ID."
)

func Connect() {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")

	if err != nil {
		panic(err)
	}

	sqlStatement := `
		create table receipts(id TEXT PRIMARY KEY, points TEXT)
	`

	_, err = db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func InsertReceipt(points string) string {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")

	if err != nil {
		panic(err)
	}

	sqlStatement := `
		insert into receipts (id, points)
		values ($1, $2)
	`

	newUUID := uuid.New()

	_, err = db.Exec(sqlStatement, newUUID, points)
	if err != nil {
		panic(err)
	}

	return newUUID.String()
}

func GetReceipts(id string) (int, error) {
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")

	if err != nil {
		panic(err)
	}

	sqlStatement := `
		select points from receipts where id=$1;
	`

	rows, err := db.Query(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return -1, errors.New(invalidIdError)
	}

	var number int
	if err := rows.Scan(&number); err != nil {
		panic(err)
	}
	return number, nil
}
