package main

import (
	"autosql/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	data := database.Database{User:"root", Password:"root", Addr:"127.0.0.1", Port:"3306", Db:"ay_wisdom", Prefix:"ay_"}
	database.Init(data)
}