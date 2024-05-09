package author

import (
	"restAPI/db"
)

type Author struct {
	ID         int
	Age        int
	Name       string
	Theme      string
	Experience int
}

func GetAuthorById(id int) (Author, error) {
	query := "SELECT * FROM authors WHERE id = ?"

	row := db.DB.QueryRow(query, id)

	var author Author
	err := row.Scan(&author.ID, &author.Age, &author.Theme, &author.Name, &author.Experience)

	if err != nil {
		return Author{}, err
	}

	return author, nil
}

func GetAuthors() ([]Author, error) {
	query := "SELECT * FROM authors"

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authors := []Author{}

	for rows.Next() {
		var author Author
		err := rows.Scan(&author.ID, &author.Age, &author.Theme, &author.Name, &author.Experience) // Data write into fields
		// by using pointers

		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	return authors, nil
}

func (a Author) PostAuthor() error {
	query := `
	INSERT INTO authors(id, age, name, theme, experience)
	VALUES(?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.ID, a.Age, a.Name, a.Theme, a.Experience)

	if err != nil {
		return err
	}

	return err
}

func (a Author) UpdateAuthor() error {
	query := `
	UPDATE authors
	SET age = ?, name = ?, theme = ?, experience = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(a.ID, a.Age, a.Name, a.Theme, a.Experience)
	return err
}

func (a Author) Delete(id int) error {
	query := "DELETE FROM authors WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
