package main

import (
	"fmt"
	"sqlib/sqlib"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Datenbank wird erfolgreich erstellt --------------------------

	err := sqlib.CreateNewDatabase("TEST")
	if err != nil {
		fmt.Println(err.Error())
	}

	// --------------------------------------------------------------

}
