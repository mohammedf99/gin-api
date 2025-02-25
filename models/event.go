package models

import (
	"example/goUdemyRest/db"
	"time"
)

// `binding:"required"` is used with ShouldBindJSON for required fields.
type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int64     `json:"userId"`
}

// var events = []Event{}
var events []Event = []Event{}

func (e *Event) Save() error {
	q := `
	INSERT INTO events(name, description, location, datetime, userId) 
	VALUES(?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)

	defer stmt.Close()

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	e.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	q := `SELECT * FROM events`

	rows, err := db.DB.Query(q)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var e Event
		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}

	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	q := `
		SELECT * FROM events
		WHERE id = ?
	`
	row := db.DB.QueryRow(q, id) // QueryRow returns one row at most.

	var e Event
	err := row.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserID)

	if err != nil {
		return nil, err
	}

	return &e, nil

}

func (e Event) Update() error {
	q := `
	UPDATE events
	SET name = ?, description = ?, location = ?, datetime = ?, userId = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)

	return err
}

// Course implementation

func (e Event) Delete() error {
	q := `
	DELETE FROM events WHERE id = ?
	`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

// REGISTER
func (e Event) Register(uId int64) error {
	q := `
	INSERT INTO registrations(eventId, userId)
	VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(q)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, e.UserID)

	return err

}

func (e Event) CancelRegistration(uId int64) error {
	q := `
    DELETE FROM registrations 
    WHERE eventId = ? AND userId = ?
    `

	stmt, err := db.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, uId) // Use uId instead of e.UserID

	return err

}

// My implementation

// func DeleteEventById(id int64) error {
// 	q := `
// 	 DELETE FROM events WHERE id = ?
// 	`

// 	stmt, err := db.DB.Prepare(q)

// 	if err != nil {
// 		return err
// 	}

// 	_, err = stmt.Exec(id)
// 	defer stmt.Close()

// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
