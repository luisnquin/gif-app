package user

import (
	"github.com/luisnquin/meow-app/src/server/models"
	"github.com/luisnquin/meow-app/src/server/store"
)

func Save(u models.User) error {
	query := `INSERT INTO users(
				username, firstname, lastname, email, password, role, created_at, updated_at
			) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

	result, err := store.DB.Exec(query, u.Username,
		u.Firstname, u.Lastname, u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		return err
	}

	if r, _ := result.RowsAffected(); r == 0 {
		return store.ErrFailedToSaveInDB
	}

	return nil
}

func UsernameOrEmailExists(username, email string) (bool, error) {
	var exists bool

	query := "SELECT exists(SELECT * FROM users WHERE username=$1 OR email=$2);"

	err := store.DB.QueryRow(query, username, email).Scan(&exists)
	if err != nil {
		return true, err
	}

	return exists, nil
}

func GetByEmailOrUsername(username, email string) (models.User, error) {
	var user models.User

	query := `SELECT  id, username, firstname, lastname, email, password, role, birthday, created_at, updated_at 
				FROM users WHERE username=$1 OR email = $2 LIMIT 1;`

	err := store.DB.QueryRow(query, username, email).Scan(&user.ID, &user.Username, &user.Firstname,
		&user.Lastname, &user.Email, &user.Password, &user.Role, &user.Birthday, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}

	return user, nil
}
