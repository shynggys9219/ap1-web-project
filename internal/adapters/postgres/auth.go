package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	pg "github.com/jackc/pgx/v5"
	"github.com/shynggys9219/ap1-web-project/internal/model"
)

type Auth struct {
	DB *pg.Conn
}

func NewAuth(conn *pg.Conn) *Auth {
	return &Auth{
		DB: conn,
	}
}

// Create This will create a new user in database
func (m *Auth) Create(user model.User) (int, error) {
	ctx := context.Background()
	stmt := `INSERT INTO users (email, passwordhash, createdAt, updatedAt)
VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id`

	var id int
	err := m.DB.QueryRow(ctx, stmt, user.Email, user.PasswordHash).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Get This will return a specific user based on id.
func (m *Auth) Get(id int) (*model.User, error) {
	ctx := context.Background()
	stmt := `SELECT * FROM snippets WHERE id=$1`

	user := &model.User{}
	err := m.DB.QueryRow(ctx, stmt, id).Scan(user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("snippet not found")
		}
	}

	return user, nil
}

// Update This will update a specific user based on id.
func (m *Auth) Update(id int) (*model.Snippet, error) {
	return nil, nil
}

// Delete This will delete a specific user based on id.
func (m *Auth) Delete(id int) error {
	return nil
}
