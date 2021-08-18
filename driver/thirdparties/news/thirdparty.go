package news

import (
	"context"
	"encoding/json"
	"jokibro/bussiness/news"
	"net/http"
)

type newsRepository struct {
	Host       string
	Key        string
	httpClient *http.Client
}

func NewNewsRepository(Host, Key string, httpClient *http.Client) news.Repository {
	return &newsRepository{
		Host,
		Key,
		httpClient,
	}
}

func (r *newsRepository) Find(ctx context.Context) (res []news.Domain, err error) {
	req, err := http.NewRequest("GET", r.Host+"/v2/top-headlines?country=id&apiKey="+r.Key, nil)
	if err != nil {
		return res, err
	}
	resp, err := r.httpClient.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return res, err
	}

	for _, d := range data.Articles {
		res = append(res, *d.ToDomain())
	}

	return res, nil
}
