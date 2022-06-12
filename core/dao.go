package core

type DAO[T any] interface {
	Get(id int64) (*T, error)
	GetAll() ([]T, error)
	Create(resource T) (*T, error)
	Update(id int64, resource T) error
	Delete(id int64) error
}
