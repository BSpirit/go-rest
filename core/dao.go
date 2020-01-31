package core

type DAO interface {
	Get(id int64) (interface{}, error)
	GetAll() (interface{}, error)
	Create(ressource interface{}) (interface{}, error)
	Update(id int64, ressource interface{}) error
	Delete(id int64) error
}
