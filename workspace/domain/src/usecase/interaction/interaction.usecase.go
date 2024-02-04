package interaction

import (
	"context"
	"net/http"

	"dealls.test/domain/src/data/interactions"
	"dealls.test/domain/src/data/users"
	"dealls.test/material/src/client"
	"dealls.test/material/src/contract"
	"dealls.test/material/src/core"
)

type Usecase interface {
	GetUserInteraction(ctx context.Context, params ParamGetUserInteraction) ([]users.EntityModel, *core.Error)
	SaveInteractionAction(ctx context.Context, params ParamSaveInteraction) *core.Error
}

type usecase struct {
	*client.Client
	*Repository
}

type Repository struct {
	User        users.Repository
	Interaction interactions.Repository
}

func NewService(c *client.Client, r *Repository) Usecase {
	return &usecase{
		Client:     c,
		Repository: r,
	}
}

func (uc *usecase) GetUserInteraction(ctx context.Context, params ParamGetUserInteraction) ([]users.EntityModel, *core.Error) {
	dbx := uc.Dbi.Cnx(ctx)

	user := &users.EntityModel{
		MetaEntity: contract.MetaEntity{
			ShowableEntity: contract.ShowableEntity{
				SecondaryId: params.UserID,
			},
		},
	}

	if err := uc.User.SelectBySecondaryId(dbx, user); err != nil {
		return nil, err
	}

	if !user.IsExist() {
		return nil, &core.Error{
			Message:    "users not found",
			StatusCode: http.StatusBadRequest,
		}
	}

	limit := 10
	hasPrivilages := false

	for _, v := range user.Privilages {
		if v.Feature == "quota" {
			hasPrivilages = true
			break
		}
	}

	if !hasPrivilages {
		count, err := uc.Interaction.CountInteractionPerDay(dbx, &interactions.EntityModel{
			Entity: interactions.Entity{
				OwnerID: user.ID,
			},
		})

		if err != nil {
			return nil, err
		}

		limit -= int(count)

		if limit < 0 {
			return nil, &core.Error{
				StatusCode: http.StatusGone,
				Message:    "out of limit",
			}
		}
	}

	target := []users.EntityModel{}

	if err := uc.User.FindAllUserAbleInteraction(dbx, user.ID, limit, &target); err != nil {
		return nil, err
	}

	return target, nil
}

func (uc *usecase) SaveInteractionAction(ctx context.Context, params ParamSaveInteraction) *core.Error {
	return uc.Dbi.Trx(ctx, func(tx context.Context) *core.Error {
		owner := &users.EntityModel{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: params.UserID,
				},
			},
		}

		if err := uc.User.SelectBySecondaryId(tx, owner); err != nil {
			return err
		}

		target := &users.EntityModel{
			MetaEntity: contract.MetaEntity{
				ShowableEntity: contract.ShowableEntity{
					SecondaryId: params.TargetID,
				},
			},
		}

		if err := uc.User.SelectBySecondaryId(tx, target); err != nil {
			return err
		}

		if err := uc.Interaction.InsertOrUpdateByOwnerAndTarget(tx, &interactions.EntityModel{
			Entity: interactions.Entity{
				OwnerID:  owner.ID,
				TargetID: target.ID,
				Action:   params.Action,
			},
		}); err != nil {
			return err
		}

		return nil
	})
}
