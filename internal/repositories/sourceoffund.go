package repositories

import (
	"context"
	"database/sql"
	"errors"
	"golang-e-wallet-rest-api/internal/apperrors"
	"golang-e-wallet-rest-api/internal/apperrors/errormsg"
	"golang-e-wallet-rest-api/internal/models"
)

const sourceOfFundRepo = "source of fund repository"

type SourceOfFundRepository interface {
	GetByName(ctx context.Context, sourceName *string) (*models.SourceOfFund, error)
}

type sourceOfFundRepository struct {
	dbtx dbtx
}

func NewSourceOfFundRepository(dbtx dbtx) *sourceOfFundRepository {
	return &sourceOfFundRepository{
		dbtx: dbtx,
	}
}

func (sfr *sourceOfFundRepository) GetByName(ctx context.Context, sourceName *string) (*models.SourceOfFund, error) {
	sourceOfFund := new(models.SourceOfFund)

	query := `
		SELECT
			id
		FROM
			source_of_funds
		WHERE
			name = $1
		;
	`

	row := sfr.dbtx.QueryRowContext(ctx, query, *sourceName)
	if err := row.Err(); err != nil {
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgInvalidQuery, sourceOfFundRepo)
	}

	err := row.Scan(&sourceOfFund.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperrors.NewCustomError(err, errormsg.ErrMsgSourceOfFundNotExist, sourceOfFundRepo)
		}
		return nil, apperrors.NewCustomError(err, errormsg.ErrMsgFailedToScanData, sourceOfFundRepo)
	}

	sourceOfFund.Name = *sourceName
	return sourceOfFund, nil
}
