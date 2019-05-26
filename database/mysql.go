package database

import (
	"autosql/common"
	"autosql/file"
	"database/sql"
	"fmt"
	"log"
	"strings"
	)

var db *sql.DB

type Table struct {
	Field string
	Type string
	Null string
	Key string
	Default string
	Extra string
}

type Database struct {
	User string
	Password string
	Addr string
	Port string
	Db string
	Prefix string
}


func Init(database Database)  {
	var err error
	sourceName := database.User + ":" + database.Password + "@tcp(" + database.Addr + ":" + database.Port + ")/" + database.Db + "?charset=utf8"
	db, err = sql.Open("mysql", sourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	prefix := database.Prefix
	length := len(prefix)
	tables := GetTables()
	//fmt.Println(tables)

	fields := make([][]string, 100)
	for _,table := range tables {
		if table == "" {
			continue
		}
		lengths := len(table)

		fields = GetFields(table)
		table = table[length:lengths]
		fmt.Println(table)
		path := "./model"
		filename := table+".go"
		text := "package model\n\n"
		text += "type " + common.CamelString(table) + " struct {\n"
		for _, field := range fields {
			//fileText,err := file.ReadAll("/sql/gorm.go")
			primaryKey := ""
			if err != nil {
				log.Fatal(err)
			}
			//fileContent := string(fileText)
			if field == nil {
				break
			}

			if len(field) == 0 {
				break
			}
			//fmt.Println(field)
			field[0] = strings.ReplaceAll(field[0], "\r\n", "")
			text += common.CamelString(field[0]) + " int "

			if field[3] == "PRI" {
				text += "`gorm:\"primary_key;"
				primaryKey = field[0]
			}

			if field[5] == "auto_increment" {
				text += "AUTO_INCREMENT\" "
			}

			if(primaryKey != ""){
				text += "json:\""+field[0]+"\"`\n"
			}else{
				text += "`json:\""+field[0]+"\"`\n"
			}


		}
		text += "}\n"

		fileText,err := file.ReadAll("./sql/gorm.txt")
		if err != nil {
			log.Fatal(err)
		}
		content := string(fileText)
		content = strings.Replace(content, "table", table, -1)
		content = strings.Replace(content, "Table", common.CamelString(table), -1)
		text += content
		file.CreateModelFile(path, filename, "gorm", text)

	}
}

func GetTables() []string {
	sqlQuery := "show tables"
	rows,err := db.Query(sqlQuery)

	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	list := make([]string, 100)
	j := 0

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}

		var value string
		for _, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			list[j] = value
		}

		j++

	}

	return list
}

func GetFields(table string) [][]string {
	list := make([][]string, 100)
	if table == "" {
		return list
	}
	querySql := "desc "+table
	rows,err := db.Query(querySql)

	columns, err := rows.Columns()

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	j := 0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			fmt.Println("log:", err)
			panic(err.Error())
		}

		var value string
		for i, col := range values {

			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			list[j] = append(list[j], value)
			list[j][i] = value
		}

		j++
	}

	return list
}
