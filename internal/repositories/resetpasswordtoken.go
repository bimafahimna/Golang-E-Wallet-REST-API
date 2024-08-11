package repositories

import (
	"context"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/models"
)

const resetPwdTokenRepo = "reset password token repository"

type ResetPasswordTokenRepository interface {
	SetupByUserId(ctx context.Context, userId *int64) error
	SaveTokenByUserId(ctx context.Context, token *string, userId *int64) error
	GetByUserId(ctx context.Context, userId *int64) (*models.ResetPasswordToken, error)
	ResetToken(ctx context.Context, userId *int64) error
}

type resetPasswordTokenRepository struct {
	dbtx dbtx
}

func NewResetPasswordTokenRepository(dbtx dbtx) *resetPasswordTokenRepository {
	return &resetPasswordTokenRepository{
		dbtx: dbtx,
	}
}

func (rp *resetPasswordTokenRepository) SetupByUserId(ctx context.Context, userId *int64) error {
	query := `
		INSERT INTO reset_password_tokens
			(user_id)
		VALUES
			($1)
	`
	row := rp.dbtx.QueryRowContext(ctx, query, userId)
	if err := row.Err(); err != nil {
		return apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, resetPwdTokenRepo)
	}

	return nil
}

func (rp *resetPasswordTokenRepository) SaveTokenByUserId(ctx context.Context, token *string, userId *int64) error {
	query := `
	UPDATE 
	reset_password_tokens
	SET
	token = $1,
	expired_at = NOW()+ interval '30 minutes'
	WHERE
	user_id = $2
	`

	_, err := rp.dbtx.ExecContext(ctx, query, token, userId)
	if err != nil {
		return apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, resetPwdTokenRepo)
	}

	return nil
}

func (rp *resetPasswordTokenRepository) GetByUserId(ctx context.Context, userId *int64) (*models.ResetPasswordToken, error) {
	resetPwd := new(models.ResetPasswordToken)

	query := `
		SELECT
			id,
			token,
			CASE 
                WHEN expired_at > NOW() THEN TRUE  
                ELSE  FALSE
				END is_valid,
				expired_at
				FROM
				reset_password_tokens
				WHERE
				user_id = $1
				;
				`

	err := rp.dbtx.QueryRowContext(ctx, query, userId).Scan(
		&resetPwd.Id,
		&resetPwd.Token,
		&resetPwd.IsValid,
		&resetPwd.ExpiredAt,
	)
	if err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, resetPwdTokenRepo)
	}

	return resetPwd, nil
}

func (rp *resetPasswordTokenRepository) ResetToken(ctx context.Context, userId *int64) error {
	query := `
		UPDATE 
			reset_password_tokens
		SET
			token = '',
			expired_at = NOW()
		WHERE
			user_id = $1
	`
	row := rp.dbtx.QueryRowContext(ctx, query, userId)
	if err := row.Err(); err != nil {
		return apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, resetPwdTokenRepo)
	}

	return nil
}
