package models

import (
	"errors"
	"example.com/events-go/db"
	"example.com/events-go/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO users (email, password) 
    VALUES (?, ?)
  `)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	u.Id = id

	return err
}

func (u *User) ValidateCredentials() error {
	row := db.DB.QueryRow("SELECT password FROM users WHERE email = ?", u.Email)

	var retrievePassword string

	err := row.Scan(&retrievePassword)
	if err != nil {
		return errors.New("Invalid password")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievePassword)

	if !isPasswordValid {
		return errors.New("Invalid password")
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	// this query is not prepared because its easier for the engine
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
