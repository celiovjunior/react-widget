package models

import "github.com/celiovjunior/goserverwidget/db"

func Get(id int64) (feedback Feedback, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM feedbacks WHERE id=$1`, id)

	err = row.Scan(&feedback.ID, &feedback.Type, &feedback.Comment, &feedback.Screenshot)

	return
}
