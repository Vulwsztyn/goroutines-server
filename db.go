package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type SqlInterface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
}

type SqlConnectionInterface interface {
	Open(driverName, dataSourceName string) (*sql.DB, error)
}

type DbInterface interface {
	InsertTs(id int)
	GetEntriesForRunner(id int) []Entry
}
type Db struct {
	postgres SqlInterface
}

func NewDb(sqlConnection func(driverName, dataSourceName string) (SqlInterface, error)) *Db {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		"root",
		"password",
		"postgres",
		"5432",
		"postgres")
	postgres, err := sqlConnection("postgres", url)
	if err != nil {
		panic(err)
	}
	return &Db{postgres}
}
func (this *Db) InsertTs(id int) {
	_, err := this.postgres.Exec("insert into timestamps (runner_id) values ($1)", id)
	if err != nil {
		panic(err)
	}
}

type Entry struct {
	Id        int       `json:"id"`
	RunnerId  int       `json:"runner_id"`
	Timestamp time.Time `json:"timestamp"`
}

func (this *Db) GetEntriesForRunner(id int) []Entry {
	rows, err := this.postgres.Query("SELECT * FROM timestamps WHERE runner_id = $1", id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	entries := []Entry{}
	for rows.Next() {
		var ts time.Time
		var runnerId int
		var id int
		err := rows.Scan(&id, &runnerId, &ts)
		if err != nil {
			panic(err)
		}
		entry := Entry{id, runnerId, ts}
		entries = append(entries, entry)
	}
	fmt.Println(entries)
	return entries
}
