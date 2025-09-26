package rss

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/fox998/gator/internal/database"
	"github.com/fox998/gator/internal/state"
	"github.com/lib/pq"
)

func ParsePublicationDate(date string) (time.Time, error) {

	expectedFormats := []string{
		time.Layout,
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.RFC850,
		time.RFC1123,
		time.RFC1123Z,
		time.RFC3339,
		time.RFC3339Nano,
		time.DateTime,
		time.DateOnly,
	}

	for _, format := range expectedFormats {
		t, err := time.Parse(format, date)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unexpected time format %s", date)
}

func isUniqueViolation(err error) bool {
	var pqErr *pq.Error
	return errors.As(err, &pqErr) && pqErr.Code.Name() == "unique_violation"
}

func ScrapeFeeds(st *state.State) error {
	feed, err := st.Db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get next feed to fetch. %w", err)
	}

	err = st.Db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		return fmt.Errorf("failed to mark as fetched: %w", err)
	}

	rssFeed, err := FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return fmt.Errorf("failed to fetch: %w", err)
	}

	fmt.Println("Fetched: " + rssFeed.Channel.Title)

	for _, item := range rssFeed.Channel.Item {

		publication_date, err := ParsePublicationDate(item.PubDate)
		if err != nil {
			log.Println("unexpected time format" + item.PubDate)
		}

		_, err = st.Db.CreatePost(context.Background(), database.CreatePostParams{
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: sql.NullTime{
				Time:  publication_date,
				Valid: err != nil,
			},
			FeedID: feed.ID,
		})

		if err != nil && !isUniqueViolation(err) {
			log.Println("Failed to create post" + err.Error())
		}
	}

	return nil
}
