package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/adamcheaib/blog-aggregator-golang/internal/config"
	"github.com/adamcheaib/blog-aggregator-golang/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func main() {
	dbUrl := "postgres://postgres:adamishere4you@localhost:5433/gator?sslmode=disable"
	State := &state{}
	if configuration, err := config.Read(); err != nil {
		fmt.Println("configuration file could not be read. Terimnating")
		os.Exit(1)
	} else {
		State.configuration = &configuration
	}

	commandsMap := commands{make(map[string]func(*state, command) error)}

	if len(os.Args) < 2 {
		fmt.Println("Command missing! Try something like: 'gator login [username]'")
		os.Exit(1)
	}

	commandsMap.register(State, "login", handlerLogin)
	commandsMap.register(State, "register", handlerRegister)
	commandsMap.register(State, "reset", resetDatabase)
	commandsMap.register(State, "users", listAllUsers)

	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		panic(err)
	}

	dbQueries := database.New(db)
	State.db = dbQueries

	commandsMap.run(State, command{os.Args[1], os.Args[2:]})
}
