package news

import (
	"context"
	"time"
)

type newsUsecase struct {
	newsRespository Repository
	contextTimeout  time.Duration
}

func NewNewsUsecase(timeout time.Duration, r Repository) Usecase {
	return &newsUsecase{
		contextTimeout:  timeout,
		newsRespository: r,
	}
}

func (uc *newsUsecase) Find(ctx context.Context) ([]Domain, error) {
	resp, err := uc.newsRespository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return resp, nil
}
