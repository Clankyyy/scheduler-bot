package pgstorage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type PGstorage struct {
	db *sql.DB
}

func NewPGStore() *PGstorage {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return &PGstorage{
		db: db,
	}
}

func (pgs PGstorage) AddUser(id int64, slug string) error {
	return nil
}

func (pgs PGstorage) GetSlug(id int64) (string, error) {
	return "", nil
}
