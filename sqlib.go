package sqlib

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

type sqlib struct{}

func CreateNewDatabase(filename string) error {

	if checkifexists(filename) {
		return errors.New("Datenbankdatei existiert schon ... ! Datenbank konnte nicht erstellt werden !")
	} else {
		db, err := sql.Open("sqlite3", filename)
		if err != nil {
			return errors.New("Fehler beim Erstellen der Datenbank ... !")
		}
		defer db.Close()
	}

	return nil
}

func CreateTable(database string, table string, rowlist TableRows) error {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		return errors.New("Fehler beim Öffnen der Datenbank ... !")
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)")

	return nil
}

func checkifexists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}
