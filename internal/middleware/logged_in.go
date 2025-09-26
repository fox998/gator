package middleware

import (
	"context"
	"fmt"

	"github.com/fox998/gator/internal/state"
)

func MiddlewareLoggedIn(handler func(*state.State, []string) error) func(*state.State, []string) error {

	return func(st *state.State, args []string) error {

		_, err := st.Db.GetUser(context.Background(), st.Cnf.CurrentUserName)
		if err != nil {
			fmt.Println(err.Error())
			return fmt.Errorf("failed to validate current user")
		}

		return handler(st, args)
	}
}
