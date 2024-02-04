package interactions

import (
	"context"
	"net/http"

	"dealls.test/material/src/client/postgresql"
	"dealls.test/material/src/core"
)

type Repository interface {
	CountInteractionPerDay(ctx context.Context, model *EntityModel) (int64, *core.Error)
	InsertOrUpdateByOwnerAndTarget(ctx context.Context, model *EntityModel) *core.Error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) CountInteractionPerDay(ctx context.Context, model *EntityModel) (int64, *core.Error) {
	var count int64
	dbi := ctx.(*postgresql.Connection).Conn

	err := dbi.
		Where("owner_id = ?", model.OwnerID).
		Where("TO_CHAR(updated_at , 'ddmmyyyy') = TO_CHAR(now(), 'ddmmyyyy')").
		Find(&model).
		Count(&count).
		Error

	if err != nil {
		return 0, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Message:    "Failed to get interaction per day",
			Origin:     err,
		}
	}

	return count, nil
}

func (r *repository) InsertOrUpdateByOwnerAndTarget(ctx context.Context, model *EntityModel) *core.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	tx := dbi.Model(model).
		Where("target_id = ?", model.TargetID).
		Where("owner_id = ?", model.OwnerID).
		Update("action", model.Action)

	if err := tx.Error; err != nil {
		return &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to insert interaction",
		}
	}

	if tx.RowsAffected == 0 {
		if err := r.Insert(ctx, model); err != nil {
			return err
		}
	}

	return nil
}

func (r *repository) Insert(ctx context.Context, model *EntityModel) *core.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	if err := dbi.Create(model).Error; err != nil {
		return &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to insert interaction",
		}
	}

	return nil
}
