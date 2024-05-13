package models

import (
	"restAPI/db"
)

type Events struct {
	ID             int
	TotalPeople    int    `binding:"required"`
	Theme          string `binding:"required"`
	MinuteDuration int
	UserId         int
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
		err := rows.Scan(&event.ID, &event.TotalPeople, &event.Theme, &event.MinuteDuration, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func (e *Events) Save() error {
	query := `
	INSERT INTO events(TotalPeople, Theme, MinuteDuration, user_id)
	VALUES(?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.TotalPeople, e.Theme, e.MinuteDuration, e.UserId)

	if err != nil {
		return err
	}

	EventId, err := result.LastInsertId()

	e.ID = int(EventId)

	return err
}

func GetEventById(id int) (Events, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Events
	err := row.Scan(&event.ID, &event.TotalPeople, &event.Theme, &event.MinuteDuration, &event.UserId)

	if err != nil {
		return Events{}, err
	}

	return event, nil
}

func (event Events) Update() error {
	query := `
	UPDATE events
	SET TotalPeople = ?, Theme = ?, MinuteDuration = ?, user_id = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.TotalPeople, event.Theme, event.MinuteDuration, event.UserId, event.ID)
	return err
}

func (e Events) Delete(id int) error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
