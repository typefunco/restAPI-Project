package models

import (
	"errors"
	"restAPI/db"
	"restAPI/utils"
)

type User struct {
	Id       int
	Login    string
	Password string
}

func (u User) Save() error {
	query := `
	INSERT INTO users (login, password)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	if len(u.Password) < 8 {
		return errors.New("SHORT PASSWORD")
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Login, hashedPassword)

	if err != nil {
		return err
	}

	UserId, err := result.LastInsertId()

	u.Id = int(UserId)

	return err

}
