package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/adamcheaib/blog-aggregator-golang/internal/database"
	"github.com/google/uuid"
	"log"
	"os"
	"time"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		log.Fatal("missing argument for the login command")
	}

	name := cmd.args[0]

	row, err := s.db.GetUser(context.Background(), name)

	if err != nil {
		fmt.Println(fmt.Sprintf("Could not find %q in the database", name))
		return err
	}

	err = s.configuration.SetUser(name)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%v has been signed in!", row.Name))

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		fmt.Println("no name passed for the register")
		return errors.New("missing argument for register")
	}

	name := cmd.args[0]

	newUser := database.CreateuserParams{
		uuid.New(),
		time.Now(),
		time.Now(),
		name,
	}

	_, err := s.db.Createuser(context.Background(), newUser)

	if err != nil {
		fmt.Println("Username already taken")
		os.Exit(1)
	}

	err = s.configuration.SetUser(name)
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("%q has been registerd", name))
	return nil
}

func resetDatabase(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		fmt.Println("Something went wrong with resetting the database!")
		return err
	}

	fmt.Println("Database has been reset")
	os.Exit(0)
	return nil
}

func listAllUsers(s *state, cmd command) error {
	rows, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Println("There was an error getting all the users")
		os.Exit(1)
	}

	for _, name := range rows {
		output := "- "

		if name == s.configuration.Current_user_name {
			output += fmt.Sprintf("'%v (current)'", name)
		} else {
			output += fmt.Sprintf("'%v'", name)
		}

		fmt.Println(output)
	}

	return nil
}
