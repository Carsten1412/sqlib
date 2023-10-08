package main

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// Datenbank wird erfolgreich erstellt --------------------------

	// err := sqlib.CreateNewDatabase("TEST")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// --------------------------------------------------------------

	// Table erstellen und Zeilen einfügen --------------------------

	// rowlist := sqlib.NewRowsList() // zuerst muss der Datentyp Tablerows erstellt werde, um die Row für die Datenbank bereitzustellen
	// rowlist.AddRow("name", "TEXT") // hier werden die einzelnen Zeilen hinzugefügt
	// rowlist.AddRow("vorname", "TEXT")
	// rowlist.AddRow("age", "INTEGER")

	// err := sqlib.CreateTable("TEST", "Person", rowlist) // hier wir der Table erstellt und die vordefnierten Rows hinzugefügt
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	//--------------------------------------------------------------
}
