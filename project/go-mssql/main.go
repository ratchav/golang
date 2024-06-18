package main

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func main() {
	connString := "sqlserver://DESKTOP-334VGBG:1433?database=gorm&integratedSecurity=true"

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error creating connection pool: ", err.Error())
		return
	}

	defer db.Close()
	defer callFuncWhenEnd()
	err = db.Ping()
	if err != nil {
		fmt.Println("Error pinging database: ", err.Error())
		return
	}

	fmt.Println("Connected!")
}

func callFuncWhenEnd() {
	fmt.Print("End")
}
