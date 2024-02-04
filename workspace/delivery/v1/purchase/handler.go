package purchase

import (
	"net/http"

	"dealls.test/domain/src/usecase/purchase"
	"dealls.test/material/src/contract"
	"dealls.test/material/src/core"
)

func (h *handler) PURCHASE(c core.Context) *core.Error {
	ctx := c.Context()
	user := c.User()

	req := PurchasePrivilagesRequest{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := c.Validate(req); err != nil {
		return err
	}

	if !req.Feature.IsValid() {
		return &core.Error{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid feature type",
		}
	}

	if err := h.uc.Purchase.PurchasePrivilages(ctx, &purchase.ParamPurchasePrivilages{
		UserID:  user.SecondaryId,
		Feature: string(req.Feature),
	}); err != nil {
		return &core.Error{}
	}

	return c.JSON(http.StatusCreated, PurchasePrivilagesResponse{
		ResponseMeta: contract.ResponseMeta{
			Message: "purchase success",
		},
	})
}
