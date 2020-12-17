package lib

import (
	"fmt"
	"net/http"

	loghttp "github.com/motemen/go-loghttp"
	"github.com/robtec/newsapi/api"
)

// GetNews uses the NewsAPI to get the latest news.
func GetNews(apiKey string, cnt int64) ([]string, error) {
	httpClient := &http.Client{
		Transport: &loghttp.Transport{},
	}
	url := "https://newsapi.org"

	client, err := api.New(httpClient, apiKey, url)
	if err != nil {
		return nil, err
	}

	opts := api.Options{
		Language: "en",
		Sources:  "new-york-times,associated-press,the-washington-post",
	}
	resp, err := client.TopHeadlines(opts)
	if err != nil {
		return nil, err
	}

	var ret []string
	for i := 0; i < 3; i++ {
		a := resp.Articles[i]
		ret = append(ret, fmt.Sprintf("%q - %s", a.Title, a.Source.Name))
	}
	return ret, nil
}
