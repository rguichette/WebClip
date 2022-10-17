package mysql

import (
	"database/sql"
	"errors"

	"github.com/rguichette/webclip/pkg/models"
)

type ClipModel struct {
	DB *sql.DB
}

func (m *ClipModel) Insert(title, content, expires string) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}
	id, err := result.LastInsertId()
	if err != nil {
		println("error: ", err)
		return 0, err
	}
	println("hellow~:", id)
	return int(id), nil

}
func (m *ClipModel) Get(id int) (*models.Clip, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)
	s := &models.Clip{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Expires, &s.Created)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

//This will return the 10 most recent
func (m *ClipModel) Latest() ([]*models.Clip, error) {
	//Write the Sql sttement we want to exec.
	stmt := `SELECT id, title, content,created,expires FROM clips WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	// use the Query() method on the connection pool to exec our SQL statement. This returns a sqlRows result set containiing the result of our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// defer to make sure sql.rows is always properly closed before Latest() returns.
	defer rows.Close()

	clips := []*models.Clip{}
	for rows.Next() {
		s := &models.Clip{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Expires, &s.Created)
		if err != nil {
			return nil, err
		}
		clips = append(clips, s)

	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return clips, nil
}

// `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`
