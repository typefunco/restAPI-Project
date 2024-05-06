package events

import (
	"restAPI/db"
)

type Events struct {
	ID             int
	TotalPeople    int    `binding:"required"`
	Theme          string `binding:"required"`
	MinuteDuration int
}

func GetEvents() ([]Events, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []Events{}

	for rows.Next() {
		var event Events
		err := rows.Scan(&event.ID, &event.TotalPeople, &event.Theme, &event.MinuteDuration)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (e Events) Save() error {
	query := `
	INSERT INTO events(id, TotalPeople, Theme, MinuteDuration)
	VALUES(?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(e.ID, e.TotalPeople, e.Theme, e.MinuteDuration)

	if err != nil {
		return err
	}

	return err
}

func GetEventById(id int) (Events, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Events
	err := row.Scan(&event.ID, &event.TotalPeople, &event.Theme, &event.MinuteDuration)

	if err != nil {
		return Events{}, err
	}

	return event, nil
}

func (event Events) Update() error {
	query := `
	UPDATE events
	SET TotalPeople = ?, Theme = ?, MinuteDuration = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.TotalPeople, event.Theme, event.MinuteDuration, event.ID)
	return err
}
