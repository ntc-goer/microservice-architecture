package repository

type RepositoryI interface {
}
type Repository struct{}

func NewRepository() *Repository {
	return &Repository{}
}
