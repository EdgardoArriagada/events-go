package models

import (
	"example.com/events-go/db"
	"time"
)

type Event struct {
	Id          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	stmt, err := db.DB.Prepare(`
    INSERT INTO events (name, description, location, date_time, user_id) 
    VALUES (?, ?, ?, ?, ?)
  `)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	// this query is not prepared because its easier for the engine
	rows, err := db.DB.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e *Event) Update() error {
	stmt, err := db.DB.Prepare(`
    UPDATE events SET name = ?, description = ?, location = ?, date_time = ? WHERE id = ?
  `)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.Id)

	if err != nil {
		return err
	}

	return nil
}

func (e *Event) Delete() error {
	stmt, err := db.DB.Prepare("DELETE FROM events WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Id)

	return err
}
