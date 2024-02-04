package privilages

import (
	"context"

	"dealls.test/material/src/client/postgresql"
	"dealls.test/material/src/core"
)

type Repository interface {
	Insert(ctx context.Context, model *EntityModel) *core.Error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Insert(ctx context.Context, model *EntityModel) *core.Error {
	dbi := ctx.(*postgresql.Connection).Conn

	if err := dbi.Create(model).Error; err != nil {
		return &core.Error{}
	}

	return nil
}
