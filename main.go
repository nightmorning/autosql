package main

import (
	"autosql/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Init()


}