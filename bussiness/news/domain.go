package news

import (
	"context"
	"time"
)

type Domain struct {
	Source      *Source
	Author      string
	Title       string
	Description string
	URL         string
	URLToImage  string
	PublishedAt time.Time
	Content     string
}

type Source struct {
	Name string
}

type Usecase interface {
	Find(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context) ([]Domain, error)
}
