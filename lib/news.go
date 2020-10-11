package lib

import (
	"fmt"
	"net/http"

	"github.com/robtec/newsapi/api"
)

func GetNews(cnt int64) ([]string, error) {
	httpClient := http.Client{}
	key := "my-api-key"
	url := "https://newsapi.org"

	client, err := api.New(httpClient, key, url)
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
