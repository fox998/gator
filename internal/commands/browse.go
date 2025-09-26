package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fox998/gator/internal/database"
	"github.com/fox998/gator/internal/state"
)

func getBrowsLimit(args []string) (uint, error) {

	size := len(args)
	if size == 1 {
		return 2, nil
	}

	if size != 2 {
		return 0, fmt.Errorf("expected zero or one param: [posts_count]")
	}

	limitStr := args[1]
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return 0, fmt.Errorf("failed convert %s to int", limitStr)
	}

	if limit < 1 {
		return 0, fmt.Errorf("post_count should be >= 1")
	}

	return uint(limit), nil
}

func Browse(st *state.State, args []string) error {
	limit, err := getBrowsLimit(args)
	if err != nil {
		return err
	}

	posts, err := st.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		Name:  st.Cnf.CurrentUserName,
		Limit: int32(limit),
	})

	if err != nil {
		return fmt.Errorf("failed to get posts for user %s: %w", st.Cnf.CurrentUserName, err)
	}

	for _, post := range posts {
		fmt.Println("---")
		fmt.Println(post.Title)
		fmt.Println(post.Url)
		fmt.Println(post.Description)
		if post.PublishedAt.Valid {
			fmt.Println("Published at:" + post.PublishedAt.Time.String())
		}
		fmt.Println("---")
	}

	return nil
}
