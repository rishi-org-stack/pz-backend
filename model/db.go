package model

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type DB struct {
	Db *pg.DB
	T  interface{}
}

func (db *DB) ConnectTodb() *DB {
	opts := &pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "rishijha",
		Database: "just",
	}
	Database := pg.Connect(opts)
	db.Db = Database
	return db
}
func (db *DB) CreateaTable() error {
	err := db.Db.Model(db.T).
		CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
	return err
}
