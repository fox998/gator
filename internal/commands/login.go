package commands

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/fox998/gator/internal/state"
)

func Login(st *state.State, args []string) error {
	if len(args) != 2 {
		return errors.New("usage: login <username>")
	}

	username := args[1]
	_, err := st.Db.GetUser(context.Background(), username)
	if err != nil {
		log.Println(err.Error())
		return errors.New("user not found")
	}

	err = st.Cnf.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}
