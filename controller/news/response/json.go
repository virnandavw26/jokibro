package response

import (
	"jokibro/bussiness/news"
)

type News struct {
	Source      *Source `json:"source"`
	Author      string  `json:"author"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	URLToImage  string  `json:"url_to_image"`
	PublishedAt string  `json:"published_at"`
	Content     string  `json:"content"`
}

type Source struct {
	Name string `json:"name"`
}

func FromDomain(domain *news.Domain) (res *News) {
	if domain != nil {
		res = &News{
			Source:      fromSourceDomain(domain.Source),
			Author:      domain.Author,
			Title:       domain.Title,
			Description: domain.Description,
			URL:         domain.URL,
			URLToImage:  domain.URLToImage,
			Content:     domain.Content,
			PublishedAt: domain.PublishedAt.UTC().Format("2006-01-02 15:04:05"),
		}
	}

	return res
}

func fromSourceDomain(domain *news.Source) (res *Source) {
	if domain != nil {
		res = &Source{
			Name: domain.Name,
		}
	}

	return res
}
