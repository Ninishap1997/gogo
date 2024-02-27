package repo

type Authorization interface {
}

type Product interface {
}

type Repo struct {
	Authorization
	Product
}

func NewRepo() *Repo {
	return &Repo{}
}
