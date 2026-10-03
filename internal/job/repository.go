package job

import (
	"context"
	"database/sql"
	"log/slog"

	commonerrors "github.com/adrian-kurek/Job-Processing-Platform/common/errors"
)

type Repository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewRepository(db *sql.DB, logger *slog.Logger) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func (jr *Repository) handleCloseErr(stmt *sql.Stmt) {
	if closeErr := stmt.Close(); closeErr != nil {
		jr.logger.Error(commonerrors.FailedToCloseStatement, "err", closeErr)
	}
}

func (jr *Repository) Insert(ctx context.Context, job CreateDTO) error {
	query := `INSERT INTO jobs (type,payload) VALUES ($1,$2)`

	stmt, err := jr.db.PrepareContext(ctx, query)
	if err != nil {
		jr.logger.Error(commonerrors.FailedToPrepareQuery, "err", map[string]string{
			"query": query,
			"error": err.Error(),
		})
		return err
	}
	defer jr.handleCloseErr(stmt)

	_, err = stmt.ExecContext(ctx, job.Type, job.Payload)
	if err != nil {
		jr.logger.Error(commonerrors.FailedToExecuteInsertQuery, "err", map[string]any{
			"query": query,
			"args": map[string]string{
				"type":    job.Type,
				"payload": job.Payload,
			},
			"error": err,
		})
		return err
	}

	return nil
}
