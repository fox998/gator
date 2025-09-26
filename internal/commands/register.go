package commands

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fox998/gator/internal/database"
	"github.com/fox998/gator/internal/state"
)

func Register(st *state.State, args []string) error {
	if len(args) != 2 {
		return errors.New("usage: register <username>")
	}

	createdUser, err := st.Db.CreateUser(context.Background(), database.CreateUserParams{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      args[1],
	})
	if err != nil {
		return err
	}

	fmt.Printf("User created ID: %v\n", createdUser.ID)
	err = st.Cnf.SetUser(createdUser.Name)
	return err
}
