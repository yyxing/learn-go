package db

import (
	"database/sql"
	"log"
)

func initConn() {
	dbConn, err := sql.Open("mysql",
		"root:root123...@tcp(154.8.173.3:33606)/video_server?charset=utf-8")
	if err != nil {
		log.Printf("init Datasource fail %s", err)
		panic(dbConn)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Printf("init Datasource fail %s", err)
		panic(err)
	}
	begin, err := dbConn.Begin()
	prepare, err := begin.Prepare("")
	prepare.Query("")
}

func selectData() {

}
