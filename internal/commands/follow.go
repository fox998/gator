package commands

import (
	"context"
	"fmt"

	"github.com/fox998/gator/internal/database"
	"github.com/fox998/gator/internal/state"
)

func Follow(st *state.State, args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("expected two params: <url_follow>")
	}

	urlToFollow := args[1]
	createdFeedFollow, err := st.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		Url:  urlToFollow,
		Name: st.Cnf.CurrentUserName,
	})
	if err != nil {
		return fmt.Errorf("failed to follow user %v: %w", urlToFollow, err)
	}

	fmt.Printf("Followed feed title: %v\n", createdFeedFollow.Title)
	fmt.Printf("Followed feed url: %v\n", createdFeedFollow.Url)
	fmt.Printf("Follower: %v\n", createdFeedFollow.FollowerName)

	return nil
}
