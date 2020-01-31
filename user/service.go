package user

import "go-rest/core"

type UserService struct {
	UserDAO        UserDAO
	UserSerializer UserSerializer
}

func (s UserService) DAO() core.DAO {
	return s.UserDAO
}

func (s UserService) Serializer() core.HTTPSerializer {
	return s.UserSerializer
}
