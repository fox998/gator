package commands

import (
	"context"
	"errors"

	"github.com/fox998/gator/internal/state"
)

func Reset(st *state.State, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: reset")
	}

	err := st.Db.Reset(context.Background())
	if err != nil {
		return errors.New("failed to reset DB: " + err.Error())
	}

	return nil
}
