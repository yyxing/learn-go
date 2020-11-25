package mapper

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Db *sql.DB
)

func init() {
	var err error
	Db, err = sql.Open("mysql",
		"root:BodenA!2019@tcp(ouroboros-mysql.mysql:30306)/ouroboros_storage?charset=utf8&parseTime=true")
	if err != nil {
		log.Printf("init Datasource fail %s", err)
		panic(Db)
	}
	err = Db.Ping()
	if err != nil {
		log.Printf("init Datasource fail %s", err)
		panic(err)
	}
}
