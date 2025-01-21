package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	pg "github.com/jackc/pgx/v5"
	"github.com/shynggys9219/ap1-web-project/internal/model"
)

// Define a SnippetModel type which wraps a sql.DB connection pool.
type Snippet struct {
	DB *pg.Conn
}

func NewSnippet(conn *pg.Conn) *Snippet {
	return &Snippet{
		DB: conn,
	}
}

// Create This will create a new snippet into the database.
func (m *Snippet) Create(title, content, expires string) (int, error) {
	ctx := context.Background()
	stmt := `INSERT INTO snippets (title, content, created, expires)
VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_DATE + $3 * INTERVAL '1 day')`

	rows, err := m.DB.Query(ctx, stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	defer rows.Close()

	log.Println(rows.Values())

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return 1, nil

}

// Get This will return a specific snippet based on its id.
func (m *Snippet) Get(id int) (*model.Snippet, error) {
	ctx := context.Background()
	stmt := `SELECT * FROM snippets WHERE id=$1`

	snippet := &model.Snippet{}
	err := m.DB.QueryRow(ctx, stmt, id).Scan(&snippet.ID, &snippet.Title, &snippet.Content, &snippet.Created, &snippet.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("snippet not found")
		}
	}

	return snippet, nil
}

// Update This will update a specific snippet based on its id.
func (m *Snippet) Update(id int) (*model.Snippet, error) {
	return nil, nil
}

// Delete This will delete a specific snippet based on its id.
func (m *Snippet) Delete(id int) (*model.Snippet, error) {
	return nil, nil
}

// Latest This will return the 10 most recently created snippets.
func (m *Snippet) Latest() ([]*model.Snippet, error) {
	ctx := context.Background()
	stmt := `SELECT * FROM snippets WHERE expires > CURRENT_DATE ORDER BY created DESC LIMIT 5`

	rows, err := m.DB.Query(ctx, stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("snippet not found")
		}
	}
	defer rows.Close()

	snippets := make([]*model.Snippet, 0)
	for rows.Next() {
		s := &model.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return snippets, nil
}
