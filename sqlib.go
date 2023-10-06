package sqlib

import (
	"database/sql"
	"errors"
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

}

func checkifexists(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil {
		return false
	} else {
		return true
	}
}
