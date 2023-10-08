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

// prüft ob die Datenbank vorhanden ist
func checkifexists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}
