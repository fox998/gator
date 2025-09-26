package commands

import (
	"context"
	"fmt"

	"github.com/fox998/gator/internal/state"
)

func Feeds(st *state.State, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("expected no params")
	}

	feeds, err := st.Db.GetFeedsWithUsername(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get feeds with username: %w", err)
	}

	for _, feed := range feeds {
		fmt.Println(feed.Title)
		fmt.Println(feed.Url)
		fmt.Println(feed.Name)
		fmt.Println()
	}

	return nil
}
