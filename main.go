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

	// err := sqlib.CreateTable("TEST", "Person", rowlist) // hier wir der Table erstellt und die vordefinierten Rows hinzugefügt
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// -------------------------------------------------------------

	//Fügt die Daten in die Rows ein ------------------------------

	// sqlib.AddValue("TEST", "Person", "name", "Mustermann2")
	// sqlib.AddValue("TEST", "Person", "vorname", "Paul2")
	// sqlib.AddValue("TEST", "Person", "age", "582")

	//--------------------------------------------------------------

	// daten aus allen Rows holen ----------------------------------
	// myMap, err := sqlib.GetAllFromTable("TEST", "Person")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(myMap)
	// -------------------------------------------------------------

	// Daten aus dem Table über einen Key holen --------------------

	// data, err := sqlib.GetFromKey("TEST", "Person", "name")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(data)

	// -------------------------------------------------------------

	// Daten aus dem Table über einen searchstring holen -----------
	// myMap, err := sqlib.Search("TEST", "Person", "name", "Mustermann")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(myMap)
	//

}
