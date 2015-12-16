package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io"
)

type dbManager struct {
	handles []*sql.DB
	pool    chan int
}

func newDbManager(dbuser, dbpassword, dbname string, poolSize int) (*dbManager, error) {

	var manager dbManager

	// setup connection login string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s", dbuser, dbpassword, dbname)

	// init the pool of connections
	manager.pool = make(chan int, poolSize)
	for i := int(0); i < poolSize; i++ {
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

func (db *dbManager) returnHandle(handle *sql.DB) {
	for i, h := range db.handles {
		if h == handle {
			db.pool <- i
			return
		}
	}
}

func (db *dbManager) getUserId(username string) (int, error) {
	// get handle on a connection
	handle := db.getHandle()
	defer db.returnHandle(handle)

	// query
	rows, err := handle.Query(`SELECT "ID" FROM "Users" WHERE "Username" = $1`, username)

	// error checks
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, io.EOF
	}

	// scan uid
	var uid int
	err = rows.Scan(&uid)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return uid, nil

}

func (db *dbManager) getUserHash(uid int) ([]byte, error) {
	// get handle on a connection
	handle := db.getHandle()
	defer db.returnHandle(handle)

	// query
	rows, err := handle.Query(`SELECT "Password" FROM "Users" WHERE "ID" = $1`, uid)

	// error checks
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, io.EOF
	}

	// scan uid
	var hash []byte
	err = rows.Scan(&hash)
	if err != nil {
		return nil, err
	}

	return hash, nil
}
