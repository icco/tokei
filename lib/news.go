package lib

import (
	"fmt"
	"net/http"

	"github.com/robtec/newsapi/api"
)

// GetNews uses the NewsAPI to get the latest news.
func GetNews(apiKey string, cnt int64) ([]string, error) {
	httpClient := http.Client{}
	url := "https://newsapi.org"

	client, err := api.New(httpClient, apiKey, url)
	if err != nil {
		return nil, err
	}

	opts := api.Options{Language: "en", Country: "us"}
	resp, err := client.TopHeadlines(opts)
	if err != nil {
		return nil, err
	}

	var ret []string
	for _, a := range resp.Articles {
		ret = append(ret, fmt.Sprintf("%s - %q", a.PublishedAt, a.Title))
	}
	return ret, nil
}
