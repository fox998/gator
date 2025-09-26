package commands

import (
	"context"
	"fmt"

	"github.com/fox998/gator/internal/database"
	"github.com/fox998/gator/internal/rss"
	"github.com/fox998/gator/internal/state"
)

func AddFeed(st *state.State, args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("two params expected: <feed_name> <feed_url>")
	}

	user, err := st.Db.GetUser(context.Background(), st.Cnf.CurrentUserName)
	if err != nil {
		return fmt.Errorf("failed to get current user %v info. %w", st.Cnf.CurrentUserName, err)
	}

	feedUrl := args[2]
	_, err = rss.FetchFeed(context.Background(), feedUrl)
	if err != nil {
		return fmt.Errorf("failed to fetch feed for %v. %w", feedUrl, err)
	}

	newFeedRow, err := st.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		Title:  args[1],
		Url:    feedUrl,
		UserID: user.ID,
	})
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("failed to create feed for %v", feedUrl)
	}

	fmt.Println(newFeedRow.ID)
	fmt.Println(newFeedRow.CreatedAt)
	fmt.Println(newFeedRow.UpdatedAt)
	fmt.Println(newFeedRow.Url)
	fmt.Println(newFeedRow.Title)
	fmt.Println(newFeedRow.UserID)

	_, err = st.Db.CreateFeedFollowIds(context.Background(), database.CreateFeedFollowIdsParams{
		FeedID: newFeedRow.ID,
		UserID: user.ID,
	})
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("failed to create feed follow. feed name %s, user %s", newFeedRow.Title, user.Name)
	}

	return nil
}
