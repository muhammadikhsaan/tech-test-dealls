package purchase

import (
	"context"
	"net/http"

	"dealls.test/domain/src/data/privilages"
	"dealls.test/domain/src/data/users"
	"dealls.test/material/src/client"
	"dealls.test/material/src/contract"
	"dealls.test/material/src/core"
)

type Usecase interface {
	PurchasePrivilages(ctx context.Context, params *ParamPurchasePrivilages) *core.Error
}

type usecase struct {
	*client.Client
	*Repository
}

type Repository struct {
	User       users.Repository
	Privilages privilages.Repository
}

func NewService(c *client.Client, r *Repository) Usecase {
	return &usecase{
		Client:     c,
		Repository: r,
	}
}

func (uc *usecase) PurchasePrivilages(ctx context.Context, params *ParamPurchasePrivilages) *core.Error {
	return uc.Dbi.Trx(ctx, func(tx context.Context) *core.Error {
		user := &users.EntityModel{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: params.UserID,
				},
			},
		}

		if err := uc.User.SelectUserIdBySecondaryId(tx, user); err != nil {
			return err
		}

		if !user.IsExist() {
			return &core.Error{
				Message:    "users not found",
				StatusCode: http.StatusBadRequest,
			}
		}

		if err := uc.Privilages.Insert(tx, &privilages.EntityModel{
			Entity: privilages.Entity{
				UserID:      user.ID,
				Feature:     params.Feature,
				ExpiredDate: nil,
			},
		}); err != nil {
			return err
		}

		return nil
	})
}
