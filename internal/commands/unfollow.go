package commands

import (
	"context"
	"fmt"

	"github.com/fox998/gator/internal/database"
	"github.com/fox998/gator/internal/state"
)

func Unfollow(st *state.State, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("one param expected")
	}

	url := args[1]
	err := st.Db.UnfollowFeed(context.Background(), database.UnfollowFeedParams{
		Url:  url,
		Name: st.Cnf.CurrentUserName,
	})

	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("failed to unfollow %s for %s", url, st.Cnf.CurrentUserName)
	}

	return nil
}
