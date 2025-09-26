package main

import (
	"log"
	"os"
	"strings"

	"github.com/fox998/gator/internal/commands"
	"github.com/fox998/gator/internal/middleware"
	"github.com/fox998/gator/internal/state"
	_ "github.com/lib/pq"
)

type commandHandler func(*state.State, []string) error

func main() {
	if len(os.Args) < 2 {
		log.Fatal("command name expected")
	}

	st := state.CreateState()

	cmds := make(map[string]commandHandler)
	cmds["login"] = commands.Login
	cmds["register"] = commands.Register
	cmds["reset"] = commands.Reset
	cmds["users"] = commands.Users
	cmds["agg"] = commands.Agg
	cmds["addfeed"] = middleware.MiddlewareLoggedIn(commands.AddFeed)
	cmds["feeds"] = commands.Feeds
	cmds["follow"] = middleware.MiddlewareLoggedIn(commands.Follow)
	cmds["following"] = middleware.MiddlewareLoggedIn(commands.Following)
	cmds["unfollow"] = middleware.MiddlewareLoggedIn(commands.Unfollow)
	cmds["browse"] = middleware.MiddlewareLoggedIn(commands.Browse)

	commandStr := strings.ToLower(os.Args[1])

	handler, found := cmds[commandStr]
	if !found {
		log.Fatal("Command not found")
	}

	err := handler(&st, os.Args[1:])
	if err != nil {
		log.Fatal(err.Error())
	}
}
