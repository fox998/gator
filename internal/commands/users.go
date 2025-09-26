package commands

import (
	"context"
	"errors"
	"fmt"

	"github.com/fox998/gator/internal/state"
)

func Users(st *state.State, args []string) error {
	if len(args) != 1 {
		return errors.New("no params expected")
	}

	users, err := st.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get users: %w", err)
	}

	for _, user := range users {
		if user == st.Cnf.CurrentUserName {
			fmt.Printf(" * %v (current)\n", user)
		} else {
			fmt.Printf(" * %v\n", user)
		}
	}

	return nil
}
