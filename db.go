package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
)

type dbManager struct {
	handles []*sql.DB
	pool    chan uint
}

func newDbManager(dbuser, dbpassword, dbname string, poolSize uint) (*dbManager, error) {

	var manager dbManager

	// setup connection login string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s", dbuser, dbpassword, dbname)

	// init the pool of connections
	for i := uint(0); i < poolSize; i++ {
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			return nil, err
		}
		manager.handles = append(manager.handles, db)
		manager.pool <- i
	}

	return &manager, nil
}

func (db *dbManager) getHandle() *sql.DB {
	i := <-db.pool
	return db.handles[i]
}

func (db *dbManager) getUserId(username string) (uint, error) {
	// get handle on a connection
	handle := db.getHandle()

	// query
	rows, err := handle.Query("SELECT id FROM users WHERE name=?", username)
	defer rows.Close()

	// error checks
	if err != nil {
		return 0, err
	}

	if !rows.Next() {
		return 0, io.EOF
	}

	// scan uid
	var uid uint
	err = rows.Scan(&uid)
	if err != nil {
		return 0, err
	}

	return uid, nil

}

/*
func (db *dbManager) getUserHash(uid uint) ([]byte, error) {

}
*/
