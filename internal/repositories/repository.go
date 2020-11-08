package repositories

import (
	"database/sql"
	"log"
)

func Rollback(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		log.Panic(err)
	}
}
