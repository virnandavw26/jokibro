package news

import (
	"jokibro/bussiness/news"
	"time"
)

type Response struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []News `json:"articles"`
}

type News struct {
	Source      *Source   `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type Source struct {
	Name string
}

func (model *News) ToDomain() (domain *news.Domain) {
	if model != nil {
		domain = &news.Domain{
			Source:      model.Source.ToDomain(),
			Author:      model.Author,
			Title:       model.Title,
			Description: model.Description,
			URL:         model.URL,
			URLToImage:  model.URLToImage,
			PublishedAt: model.PublishedAt,
			Content:     model.Content,
		}
	}
	return domain
}

func (model *Source) ToDomain() (domain *news.Source) {
	if model != nil {
		domain = &news.Source{
			Name: model.Name,
		}
	}
	return domain
}
