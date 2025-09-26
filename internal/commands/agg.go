package commands

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/fox998/gator/internal/rss"
	"github.com/fox998/gator/internal/state"
)

func Agg(st *state.State, args []string) error {
	if len(args) != 2 {
		return errors.New("no params expected: <time_between_reqs>")
	}

	intervalStr := args[1]
	fetchInterval, err := time.ParseDuration(intervalStr)
	if err != nil {
		return fmt.Errorf("failed to parse duration: %s", intervalStr)
	}

	const MinFetchInterval = 500 * time.Millisecond
	if fetchInterval < MinFetchInterval {
		return fmt.Errorf("intreval should be greater than %v", MinFetchInterval)
	}

	fmt.Printf("Collecting feeds every %v\n", fetchInterval)
	ticker := time.NewTicker(fetchInterval)
	for ; ; <-ticker.C {
		err = rss.ScrapeFeeds(st)
		if err != nil {
			log.Printf("Failed to crape feed %v\n", err.Error())
		}
	}
}
