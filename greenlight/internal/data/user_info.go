package data

import (
	"database/sql"
	"errors"
	"time"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type UserInfo struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Fname        string    `json:"fname"`
	Sname        string    `json:"sname"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"password_hash"`
	UserRole     string    `json:"user_role"`
	Activated    bool      `json:"activated"`
	Version      int       `json:"version"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) Insert(userInfo *UserInfo) error {
	query := `
        INSERT INTO user_1nfo (created_at, updated_at, fname, sname, email, password_hash, user_role, activated, version)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id
    `
	return m.DB.QueryRow(
		query,
		time.Now(),
		time.Now(),
		userInfo.Fname,
		userInfo.Sname,
		userInfo.Email,
		userInfo.PasswordHash,
		userInfo.UserRole,
		userInfo.Activated,
		userInfo.Version,
	).Scan(&userInfo.ID)
}

func (m UserModel) Get(id int64) (*UserInfo, error) {
	query := `
        SELECT * FROM user_1nfo WHERE id = $1
    `
	userInfo := &UserInfo{}
	err := m.DB.QueryRow(query, id).Scan(
		&userInfo.ID,
		&userInfo.CreatedAt,
		&userInfo.UpdatedAt,
		&userInfo.Fname,
		&userInfo.Sname,
		&userInfo.Email,
		&userInfo.PasswordHash,
		&userInfo.UserRole,
		&userInfo.Activated,
		&userInfo.Version,
	)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (m UserModel) GetAll() ([]*UserInfo, error) {
	query := `
        SELECT * FROM users_info
    `
	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*UserInfo
	for rows.Next() {
		userInfo := &UserInfo{}
		err := rows.Scan(
			&userInfo.ID,
			&userInfo.CreatedAt,
			&userInfo.UpdatedAt,
			&userInfo.Fname,
			&userInfo.Sname,
			&userInfo.Email,
			&userInfo.PasswordHash,
			&userInfo.UserRole,
			&userInfo.Activated,
			&userInfo.Version,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, userInfo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (m UserModel) Delete(id int64) error {
	query := `
        DELETE FROM users_info
        WHERE id = $1
    `
	_, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
