package main

import (
	"errors"
	"fmt"
	"os"

	"gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return handler(s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is required")
	}
	username := cmd.args[0]

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Printf("user set to %s\n", username)
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading config:", err)
		os.Exit(1)
	}

	s := &state{cfg: &cfg}

	cmds := &commands{handlers: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: gator <command> [args...]")
		os.Exit(1)
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err := cmds.run(s, cmd); err != nil {
		if !errors.Is(err, nil) {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
