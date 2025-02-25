package models

import (
	"errors"
	"example/goUdemyRest/db"
	"example/goUdemyRest/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	q := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hp, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hp)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = userId

	return err

}

func (u *User) ValidateCredentials() error {
	q := `SELECT id, password FROM users WHERE email = ?`

	row := db.DB.QueryRow(q, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	// defer db.DB.Close()

	isPasswordValid := utils.CheckPasswordHash(retrievedPassword, u.Password)

	if !isPasswordValid {
		return errors.New("wrong user credentials")
	}

	return nil
}
