package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"regexp"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FormatResultString(input string) string {
	re := regexp.MustCompile("<[^>]*>")
	stripped := re.ReplaceAllString(input, "")
	escaped := html.UnescapeString(stripped)

	return escaped
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("request creation failed for %s", feedURL)
	}

	request.Header.Set("User-Agent", "gator")

	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("client Do failed")
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("read all failed")
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(data, &feed)
	if err != nil {
		fmt.Println(err.Error())
		return nil, fmt.Errorf("unmarshal failed %s", string(data))
	}

	feed.Channel.Title = FormatResultString(feed.Channel.Title)
	feed.Channel.Description = FormatResultString(feed.Channel.Description)

	items := feed.Channel.Item
	for i, item := range items {
		items[i].Title = FormatResultString(item.Title)
		items[i].Description = FormatResultString(item.Description)
	}

	return &feed, nil
}
