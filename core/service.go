package core

type Service interface {
	DAO() DAO
	Serializer() HTTPSerializer
}
