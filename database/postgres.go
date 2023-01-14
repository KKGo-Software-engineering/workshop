package database

import (
	"database/sql"
	"log"

	"github.com/kkgo-software-engineering/workshop/pocket"
)

func createAccTable(db *sql.DB) {

	initTable := `
	CREATE TABLE IF NOT EXISTS "accounts" (
		"id" int4 NOT NULL DEFAULT nextval('account_id'::regclass),
		"balance" float8 NOT NULL DEFAULT 0,
		PRIMARY KEY ("id")
	);
	`
	_, err := db.Exec(initTable)

	if err != nil {
		log.Fatal("Having a Problem Creating new Account Table", err)
	}

}

func createPocketTable(db *sql.DB) {

	initTable := `
	CREATE TABLE IF NOT EXISTS "pockets" (
		"id" int4 NOT NULL,
		"name" TEXT NOT NULL,
		"category" TEXT NOT NULL,
		"Currency" TEXT NOT NULL,
		"balance" float8 NOT NULL DEFAULT 0,
		PRIMARY KEY ("id")
	);
	`
	_, err := db.Exec(initTable)

	if err != nil {
		log.Fatal("Having a Problem Creating new Pocket Table", err)
	}

}

func GetAllPockets(db *sql.DB, id string) ([]pocket.Pocket, error) {

	queryStatement := `
	SELECT * FROM pockets Where id = $1
	`

	st, err := db.Prepare(queryStatement)

	if err != nil {
		return nil, err
	}

	rows, err2 := st.Query(id)
	if err2 != nil {
		return nil, err
	}

	pockets := []pocket.Pocket{}

	for rows.Next() {

		pocket := pocket.Pocket{}

		err = rows.Scan(&pocket.Id, &pocket.Name, &pocket.Category, &pocket.Currency, &pocket.Balance)

		if err != nil {
			return pockets, err
		}

		pockets = append(pockets, pocket)
	}

	return pockets, nil

}

func CreateTables(db *sql.DB) {

	// createAccTable(db)
	createPocketTable(db)

}
