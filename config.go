package main

import (
	"errors"
	"fmt"
	"github.com/adamcheaib/blog-aggregator-golang/internal/config"
	"github.com/adamcheaib/blog-aggregator-golang/internal/database"
	"log"
)

type state struct {
	configuration *config.Config
	db            *database.Queries
}

type command struct {
	name string
	args []string
}

type commands struct {
	list map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	commandFunc, exists := c.list[cmd.name]
	if !exists {
		return errors.New(fmt.Sprintf("the command %v does not exist", cmd.name))
	}

	if err := commandFunc(s, cmd); err != nil {
		return err
	}

	return nil
}

func (c *commands) register(s *state, cmdName string, callback func(*state, command) error) {
	if _, exists := c.list[cmdName]; exists {
		log.Fatal(fmt.Sprintf("could not register %v-command. It already exists", cmdName))
	}

	c.list[cmdName] = callback
}
