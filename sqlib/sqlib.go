package sqlib

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// type sqlib struct{}

// erstellt eine neue Datenbank
//
// Beispiel :
//
// err := sqlib.CreateNewDatabase("TEST")
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
func CreateNewDatabase(filename string) error {

	if checkifexists(filename) {
		return errors.New(" Datenbankdatei existiert schon ... ! Datenbank konnte nicht erstellt werden")
	} else {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err.Error())
			return errors.New(" Fehler beim Erstellen der Datenbank ... ")
		}
		defer file.Close()
	}

	return nil
}

// erstellt eine neuen Table mit den mit TableRows erstellten Liste

// Beispiel :

// zuerst muss der Datentyp Tablerows erstellt werde, um die Row für die Datenbank bereitzustellen
// rowlist := sqlib.NewRowsList()
// hier werden die einzelnen Zeilen hinzugefügt

// rowlist.AddRow("name", "TEXT")
// rowlist.AddRow("vorname", "TEXT")
// rowlist.AddRow("age", "INTEGER")

// hier wir der Table erstellt und die vordefnierten Rows hinzugefügt
// err = sqlib.CreateTable("TEST", "Person", rowlist)

//	if err != nil {
//		fmt.Println(err.Error())
//	}
func CreateTable(database string, table string, rowlist TableRows) error {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return errors.New(" Fehler beim Öffnen der Datenbank ... ")
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", table, rowlist.ToString())
	fmt.Println(query)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New(" Fehler beim Ausführen des CREATE TABLE Befehls ... ")
	}

	return nil
}

// Fügt die Daten in die Rows ein
//
// sqlib.AddValue("TEST", "Person", "name", "Mustermann2")
//
// sqlib.AddValue("TEST", "Person", "vorname", "Paul2")
//
// sqlib.AddValue("TEST", "Person", "age", "582")
func AddValue(database string, table string, key string, value string) {
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ('%s')", table, key, value)

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return
	}
	defer db.Close()

	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

// holt von allen Rows die Daten
//
// myMap, err := sqlib.GetAllFromTable("TEST", "Person")
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
func GetAllFromTable(database string, table string) ([]map[string]interface{}, error) {

	query := fmt.Sprintf("SELECT * FROM %s", table)

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return nil, errors.New(" Fehler beim Öffnen der Datenbank ... ")
	}
	defer db.Close()

	rowscontent, err := db.Query(query)
	if err != nil {
		return nil, errors.New(" Fehler beim Durchsuchen der Datenbank ... ")
	}
	defer rowscontent.Close()

	columns, err := rowscontent.Columns()
	if err != nil {
		return nil, errors.New(" Fehler beim Auslesen der Spalten ... ")
	}

	values := make([]interface{}, len(columns))
	valueoptrs := make([]interface{}, len(columns))

	for i := range columns {
		valueoptrs[i] = &values[i]
	}

	data := []map[string]interface{}{}

	for rowscontent.Next() {
		err := rowscontent.Scan(valueoptrs...)
		if err != nil {
			return nil, errors.New(" Fehler beim Druchlaufen der Scanvorgangs ... ")
		}

		for i, col := range columns {
			val := values[i]
			if val != nil {
				newmap := map[string]interface{}{col: val}
				data = append(data, newmap)
			}
		}
	}

	return data, nil
}

// holt von einem bestimmten Key die Daten
//
// data, err := sqlib.GetFromKey("TEST", "Person", "name")
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
// fmt.Println(data)
func GetFromKey(database string, table string, key string) ([]map[string]interface{}, error) {

	query := fmt.Sprintf("SELECT %s FROM %s", key, table)

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return nil, errors.New(" Fehler beim Öffnen der Datenbank ... ")
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, errors.New(" Fehler beim Durchsuchen der Datenbank ... ")
	}
	defer rows.Close()

	data := make([]map[string]interface{}, 0)
	column, err := rows.Columns()

	if err != nil {
		return nil, errors.New(" Fehler beim Lesen der Columns ... ")
	}

	for rows.Next() {
		var value interface{}

		err := rows.Scan(&value)
		if err != nil {
			return nil, errors.New(" Fehler beim Auslesen der Datenbank ... ")
		}

		if value != nil {
			newMap := map[string]interface{}{column[0]: value}
			data = append(data, newMap)
		}

	}

	return data, nil
}

// sucht in einer bestimmten Spalte der Datenbank nach einem Suchbegriff
//
// myMap, err := sqlib.Search("TEST", "Person", "name", "Mustermann")
//
//	if err != nil {
//		fmt.Println(err.Error())
//	}
//
// fmt.Println(myMap)
func Search(database string, table string, column string, searchstring string) ([]map[string]interface{}, error) {

	query := fmt.Sprintf("SELECT * FROM %s WHERE %s LIKE ? ", table, column)

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return nil, errors.New(" Fehler beim Öffnen der Datenbank ... ")
	}
	defer db.Close()

	rows, err := db.Query(query, "%"+searchstring+"%")
	if err != nil {
		return nil, errors.New(" Fehler beim Durchsuchen der Datenbank ... ")
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, errors.New(" Fehler beim Auflisten der Spalten ... ")
	}

	values := make([]interface{}, len(columns))
	valuesptr := make([]interface{}, len(columns))

	data := make([]map[string]interface{}, 0)

	for i := range columns {
		valuesptr[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuesptr...)
		if err != nil {
			return nil, errors.New(" Fehler beim Scannen der Reihen ... ")
		}

		for i, v := range columns {
			val := values[i]
			if val != nil {
				newMap := map[string]interface{}{v: val}
				data = append(data, newMap)
			}
		}
	}

	return data, nil
}

// such in der Datenbank in einem bestimmten Key nach dem Suchbegriff
func searchKey() {}

// prüft ob die Datenbank vorhanden ist
func checkifexists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}
