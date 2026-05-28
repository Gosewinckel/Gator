package main

import (
	"github.com/Gosewinckel/Gator/internal/config"
	"github.com/Gosewinckel/Gator/internal/database"
)

type state struct {
	db *database.Queries
	conf *config.Config
}

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func initState() state {
	cnf := config.Read()
	s := state{}
	s.conf = &cnf
	return s
}

func initCommands(s *state) commands {
	commands := commands{}
	commandMap := make(map[string]func(*state, command) error)
	commands.handlers = commandMap
	
	// register fuunctions
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAggregator)
	commands.register("addfeed", handlerFeed)
	commands.register("feeds", handlerFeeds)
	commands.register("follow", handlerFollow)
	commands.register("following", handlerFollowing)
	return commands
}

func (c *commands) run(s *state, cmd command) error {
	fun := c.handlers[cmd.Name]
	err := fun(s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
