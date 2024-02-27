package service

import "awesomeProject/pkg/repo"

type Authorization interface {
}

type Product interface {
}

type Service struct {
	Authorization
	Product
}

func NewService(repos *repo.Repo) *Service {
	return &Service{}
}
