package user_repository

import (
	"database/sql"
	"errors"
	"github.com/fyllekanin/go-server/src/entities"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func (repository *UserRepository) DoUsernameExist(username string) (bool, error) {
	statement, err := repository.db.Prepare("SELECT COUNT(*) FROM users WHERE username = $1 LIMIT 1")
	if err != nil {
		log.Println(err)
		return false, errors.New("failed to prepare statement")
	}
	rows, err := statement.Query(username)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return false, errors.New("failed to query statement")
	}
	var count int
	for rows.Next() {
		rows.Scan(&count)
	}
	return count > 0, err
}

func (repository *UserRepository) GetUserByUsername(username string) (entities.UserEntity, error) {
	statement, err := repository.db.Prepare("SELECT * FROM users WHERE username = $1 LIMIT 1")
	if err != nil {
		log.Println(err)
		return entities.UserEntity{}, errors.New("failed to prepare statement")
	}
	rows, err := statement.Query(username)
	defer rows.Close()

	if err != nil {
		log.Println(err)
		return entities.UserEntity{}, errors.New("failed to query statement")
	}
	var response entities.UserEntity
	for rows.Next() {
		rows.Scan(&response.Id, &response.Username, &response.Password, &response.CreatedAt, &response.CreatedAt)
	}
	return response, err
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
