package repositories

import (
	"database/sql"
	"go-starter/internal/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

//UserRepository structure
type UserRepository struct {
	DbConn *sql.DB
}

func (repo UserRepository) InsertUser(user models.AddUserRequest, id uuid.UUID) (int64, error) {
	query := "INSERT INTO go_schema.user (id, name, email, mobile, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6);"

	var now = time.Now().UTC().Truncate(time.Millisecond)

	if tx, txErr := repo.DbConn.Begin(); txErr == nil {
		stmt, err := tx.Prepare(query)
		if err != nil {
			log.Println(err)
			Rollback(tx)

			return -1, err
		}
		if r, err := stmt.Exec(id, user.Name, user.Email, user.Mobile, now, now); err != nil {
			log.Println(err)
			Rollback(tx)

			return -1, err
		} else {
			id, _ := r.RowsAffected()

			err := tx.Commit()
			if err != nil {
				log.Println(err)

				return -1, err
			}

			return id, nil
		}
	} else {
		log.Println(txErr)

		return -1, txErr
	}
}
