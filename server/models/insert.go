package models

import "github.com/celiovjunior/goserverwidget/db"

func Insert(feedback Feedback) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO feedbacks (
		type, comment, screenshot
	) VALUES (
		$1, $2, $3
	) RETURNING id`

	err = conn.QueryRow(sql, feedback.Type, feedback.Comment, feedback.Screenshot).Scan(&id)

	return
}
