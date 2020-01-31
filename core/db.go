package core

import (
	"database/sql"
	"encoding/json"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(driverName string, dataSourceName string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

type NullInt64 struct {
	sql.NullInt64
}

func (v NullInt64) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Int64)
	}

	return json.Marshal(nil)
}

func (v *NullInt64) UnmarshalJSON(data []byte) error {
	var x *int64
	err := json.Unmarshal(data, &x)
	if err != nil {
		return err
	}

	if x != nil {
		v.Valid = true
		v.Int64 = *x
	} else {
		v.Valid = false
	}

	return nil
}
