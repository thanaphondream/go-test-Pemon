package service

import (
	"test/repository"
)

type ServiceView interface {
	Views() ([]repository.View, error)
}

type serviceViewRepo struct {
	repo repository.RerositoryView
}

func NewServiceView(r repository.RerositoryView) ServiceView{
	return serviceViewRepo{repo: r}
}

func (v serviceViewRepo) Views() ([]repository.View, error) {
	view, err := v.repo.View();
	if err != nil {
		return []repository.View{}, err
	}

	return view, err
}

