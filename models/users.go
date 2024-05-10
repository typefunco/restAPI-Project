package models

import (
	"errors"
	"restAPI/db"
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

	result, err := stmt.Exec(u.Login, u.Password)

	if err != nil {
		return err
	}

	UserId, err := result.LastInsertId()

	u.Id = int(UserId)

	return err

}
