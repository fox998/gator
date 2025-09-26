package state

import (
	"database/sql"
	"log"

	"github.com/fox998/gator/internal/config"
	"github.com/fox998/gator/internal/database"
)

type State struct {
	Cnf *config.Config
	Db  *database.Queries
}

func CreateState() State {
	cnf, err := config.Read()
	if err != nil {
		log.Fatal("Unable to read the config: " + err.Error())
	}

	db, err := sql.Open("postgres", cnf.DbURL)
	if err != nil {
		log.Fatal(err.Error())
	}

	return State{
		Cnf: &cnf,
		Db:  database.New(db),
	}
}
