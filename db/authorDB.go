package db

func createAuthorDB() {
	authorFields := `
	CREATE TABLE IF NOT EXISTS authors
	(id INTEGER PRIMARY KEY AUTOINCREMENT,
	age INTEGER NOT NULL,
	name TEXT NOT NULL,
	theme TEXT NOT NULL,
	experience INTEGER NOT NULL)
	`
	_, err := DB.Exec(authorFields)

	if err != nil {
		panic("Could not create Author table")
	}

}
