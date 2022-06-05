package user

import (
	"go-rest/core"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID       int64
	Username string
	Age      core.NullInt64
}

type UserDAO struct {
	DB *sqlx.DB
}

func (dao UserDAO) Get(id int64) (interface{}, error) {
	user := User{}
	err := dao.DB.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	if err != nil {
		return nil, core.Trace(err)
	}

	return user, nil
}

func (dao UserDAO) GetAll() (interface{}, error) {
	users := []User{}
	err := dao.DB.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, core.Trace(err)
	}

	return users, nil
}

func (dao UserDAO) Create(resource interface{}) (interface{}, error) {
	user := resource.(User)
	res, err := dao.DB.Exec("INSERT INTO users(username, age) VALUES(?, ?)", user.Username, user.Age)
	if err != nil {
		return nil, core.Trace(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, core.Trace(err)
	}
	user.ID = id

	return user, nil
}

func (dao UserDAO) Update(id int64, resource interface{}) error {
	user := resource.(User)
	_, err := dao.DB.Exec("UPDATE users SET username=?, age=? WHERE id=?", user.Username, user.Age, id)
	if err != nil {
		return core.Trace(err)
	}

	return nil
}

func (dao UserDAO) Delete(id int64) error {
	_, err := dao.DB.Exec("DELETE FROM users WHERE id=?", id)
	if err != nil {
		return core.Trace(err)
	}

	return nil
}
