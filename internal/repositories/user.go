package repositories

import (
	"context"
	"database/sql"
	"errors"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/models"
)

const userRepo = "user repository"

type UserRepository interface {
	SaveAccount(ctx context.Context, user *models.User) (*int64, error)
	GetByEmail(ctx context.Context, email *string) (*models.User, error)
	IncreaseGameAttempt(ctx context.Context, attempts, userId int64) error
	GetById(ctx context.Context, userId int64) (*models.User, error)
	SaveNewPassword(ctx context.Context, newPassword string, userId *int64) error
}

type userRepository struct {
	dbtx dbtx
}

func NewUserRepository(dbtx dbtx) *userRepository {
	return &userRepository{
		dbtx: dbtx,
	}
}

func (ur *userRepository) SaveAccount(ctx context.Context, user *models.User) (*int64, error) {
	query := `
		INSERT INTO users
			(username,email,password)
		VALUES
			($1,$2,$3)
		RETURNING id
	`
	row := ur.dbtx.QueryRowContext(ctx, query, user.Username, user.Email, user.Password)
	if err := row.Err(); err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgEmailExist, userRepo)
	}

	err := row.Scan(&user.Id)
	if err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToScanData, userRepo)
	}

	return &user.Id, nil
}

func (ur *userRepository) GetByEmail(ctx context.Context, email *string) (*models.User, error) {
	query := `
		SELECT
			id,
			username,
			password
		FROM
			users
		WHERE
			email ILIKE $1
		;
	`
	user := new(models.User)

	row := ur.dbtx.QueryRowContext(ctx, query, *email)
	if err := row.Err(); err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, userRepo)
	}

	err := row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewCustomError(err, errormsg.ErrMsgEmailNotExist, userRepo)
		}
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToScanData, userRepo)
	}

	user.Email = *email

	return user, nil
}

func (ur *userRepository) IncreaseGameAttempt(ctx context.Context, attempts, userId int64) error {
	query := `
	UPDATE 
		users
	SET
		game_attempts = game_attempts + $1
	WHERE
		id = $2
	`

	_, err := ur.dbtx.ExecContext(ctx, query, attempts, userId)
	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetById(ctx context.Context, userId int64) (*models.User, error) {
	user := new(models.User)

	query := `
		SELECT
			username,
			email,
			game_attempts,
			created_at,
			updated_at
		FROM
			users
		WHERE
			id = $1
	`

	err := ur.dbtx.QueryRowContext(ctx, query, userId).Scan(
		&user.Username,
		&user.Email,
		&user.GameAttempts,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToScanData, userRepo)
	}

	user.Id = userId
	return user, nil
}

func (ur *userRepository) SaveNewPassword(ctx context.Context, newPassword string, userId *int64) error {
	query := `
		UPDATE 
			users
		SET
			password = $1
		WHERE
			id = $2
	`

	_, err := ur.dbtx.ExecContext(ctx, query, newPassword, userId)
	if err != nil {
		return err
	}

	return nil
}
