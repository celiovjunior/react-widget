package models

import "github.com/celiovjunior/goserverwidget/db"

func GetAll() (feedbacks []Feedback, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM feedbacks`)

	if err != nil {
		return
	}

	for rows.Next() {
		var feedback Feedback
		err = rows.Scan(&feedback.ID, &feedback.Type, &feedback.Comment, &feedback.Screenshot)

		if err != nil {
			continue
		}

		feedbacks = append(feedbacks, feedback)
	}

	return
}
