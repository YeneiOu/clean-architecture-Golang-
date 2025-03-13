package respositories

import (
	"bank/modules/entities"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	Db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) entities.UserRepository {
	return &UserRepo{
		Db: db,
	}
}

func (r *UserRepo) GetUsers() ([]entities.UserResponse, error) {
	query := `SELECT user_id, username, password, email, user_since, last_visit FROM users`

	rows, err := r.Db.Queryx(query)
	if err != nil {
		fmt.Println(err.Error())
		return []entities.UserResponse{}, err
	}
	defer rows.Close()

	var users []entities.UserResponse

	for rows.Next() {
		var u entities.UserResponse
		if err := rows.StructScan(&u); err != nil {
			log.Println("Error scanning row:", err)
			return []entities.UserResponse{}, err
		}
		users = append(users, u)

	}

	// Ensure rows.Err() is checked
	if err := rows.Err(); err != nil {
		return []entities.UserResponse{}, err
	}

	return users, nil
}

func (r *UserRepo) GetUserByID(id int) (*entities.UserResponse, error) {
	return nil, nil
}
