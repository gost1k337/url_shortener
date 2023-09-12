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

	row := r.QueryRowContext(ctx, query, username, email, passwordHash)
	var id int64
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("scan %w", err)
	}
	return id, nil
}

func (r *UserRepo) GetById(ctx context.Context, id int64) (*entity.User, error) {
	query := `SELECT * FROM users WHERE id=$1`
	rows, err := r.QueryContext(ctx, query, id)
	if err != nil {
		if errors.Is(rows.Err(), sql.ErrNoRows) {
			return nil, repoerrors.ErrNotFound
		}
		return nil, fmt.Errorf("query: %w", err)
	}

	user := new(entity.User)

	if rows.Next() {
		if err := rows.Scan(
			&user.Id,
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
