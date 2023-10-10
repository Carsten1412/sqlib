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

// fügt Daten in die Rows ein
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
func GetFromKey() {}

// prüft ob die Datenbank vorhanden ist
func checkifexists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}
