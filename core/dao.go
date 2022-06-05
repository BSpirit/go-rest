package core

type DAO interface {
	Get(id int64) (interface{}, error)
	GetAll() (interface{}, error)
	Create(resource interface{}) (interface{}, error)
	Update(id int64, resource interface{}) error
	Delete(id int64) error
}
