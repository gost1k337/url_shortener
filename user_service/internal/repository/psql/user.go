package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/gost1k337/url_shortener/user_service/internal/entity"
	"github.com/gost1k337/url_shortener/user_service/internal/repository/repoerrors"
	"github.com/gost1k337/url_shortener/user_service/pkg/logging"
	"github.com/gost1k337/url_shortener/user_service/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
	logger logging.Logger
}

func NewUserRepo(pg *postgres.Postgres, logger logging.Logger) *UserRepo {
	return &UserRepo{
		pg,
		logger,
	}
}

func (r *UserRepo) Create(ctx context.Context, username, email, passwordHash string) (int64, error) {
	query := `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`

	row, err := r.ExecContext(ctx, query, username, email, passwordHash)
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}

	id, err := row.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("get id: %w", err)
	}

	return id, nil
}

func (r *UserRepo) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `SELECT * FROM users WHERE id=$1`

	rows, err := r.QueryContext(ctx, query, id)

	defer func() {
		err = rows.Close()
	}()

	if err != nil {
		if errors.Is(rows.Err(), sql.ErrNoRows) {
			return nil, repoerrors.ErrNotFound
		}

		return nil, fmt.Errorf("query: %w", err)
	}

	user := new(entity.User)

	if rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}
	} else {
		return nil, repoerrors.ErrNotFound
	}

	return user, nil
}

func (r *UserRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id=$1`

	_, err := r.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	return nil
}
