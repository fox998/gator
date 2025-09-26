package commands

import (
	"context"
	"fmt"

	"github.com/fox998/gator/internal/state"
)

func Following(st *state.State, args []string) error {

	if len(args) != 1 {
		return fmt.Errorf("no params expected")
	}

	feeds, err := st.Db.GetFeedFollowsForUser(context.Background(), st.Cnf.CurrentUserName)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("failed get followed feeds for %s", st.Cnf.CurrentUserName)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Title)
	}

	return nil
}
