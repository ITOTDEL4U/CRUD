package psql

import (
	"database/sql"
)

type ObjSql struct {
	db *sql.DB
}

func New(db *sql.DB) *ObjSql {
	return &ObjSql{db}
}
