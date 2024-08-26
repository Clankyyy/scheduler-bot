package pgstorage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGstorage struct {
	db *pgxpool.Pool
}

func Test() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string
	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}

func NewPGStore() *PGstorage {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	err = dbpool.Ping(context.Background())
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}

	return &PGstorage{
		db: dbpool,
	}
}

func (pgs PGstorage) CreateUsersTable() error {
	schema := `CREATE TABLE IF NOT EXISTS tg_users (
  user_id bigint NOT NULL UNIQUE,
  group_slug varchar(100) DEFAULT 'default'
);`
	_, err := pgs.db.Exec(context.Background(), schema)
	if err != nil {
		log.Print("Can't create table: ", err.Error())
		return err
	}
	return nil
}

func (pgs PGstorage) AddUser(id int64, slug string) error {
	_, err := pgs.db.Exec(context.Background(), `insert into tg_users(user_id, group_slug) values($1, $2)
	on conflict (user_id) do update set group_slug=excluded.group_slug`, id, slug)
	return err
}

func (pgs PGstorage) GetSlug(id int64) (string, error) {
	var slug string
	err := pgs.db.QueryRow(context.Background(), "select group_slug from tg_users where user_id=$1", id).Scan(&slug)
	return slug, err
}
