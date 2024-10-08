package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewRepoOfUsers(db *sql.DB) *users {
	return &users{db}
}

func (repo users) Create(user models.User) (uint64, error) {
	statement, err := repo.db.Prepare(`
		INSERT INTO users (name, nickname, email, password)
		VALUES (?, ?, ?, ?)`,
	)
	if err != nil {
		return 0, err
	}

	result, err := statement.Exec(user.Name, user.NickName, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertId), nil
}

func (repo users) Find(nameOrNickname string) ([]models.User, error) {
	nameOrNickname = fmt.Sprintf("%%%s%%", nameOrNickname)

	lines, err := repo.db.Query(`
		SELECT id, name, nickname, email, createdAt
		FROM users
		WHERE name LIKE ? OR nickname LIKE ?`,
		nameOrNickname,
		nameOrNickname,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.Id,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (repo users) FindByID(id uint64) (models.User, error) {

	line, err := repo.db.Query(`
		SELECT id, name, nickname, email, createdAt
		FROM users
		WHERE id = ?`,
		id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User

	if line.Next() {
		if err = line.Scan(
			&user.Id,
			&user.Name,
			&user.NickName,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repo users) Update(id uint64, user models.User) error {
	statement, err := repo.db.Prepare(`
		UPDATE users SET 
		name = ?, 
		nickname = ?,
		email = ?
		WHERE id = ?`,
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(
		user.Name,
		user.NickName,
		user.Email,
		id,
	); err != nil {
		return err
	}
	return nil
}
