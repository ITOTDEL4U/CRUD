package psql

import (

    "database/sql"

)

type Books struct {
    db *sql.DB
}

func NewBooks(db *sql.DB) *Books {
    return &Books{db}
}

