package purchase

import (
	"dealls.test/domain/src/usecase"
	"dealls.test/material/src/core"
)

type Handler interface {
	Router(r core.Router)
}

type handler struct {
	uc *usecase.Service
}

func NewHandler(uc *usecase.Service) Handler {
	return &handler{
		uc: uc,
	}
}

func (h *handler) Router(r core.Router) {
	r.Post("/", h.PURCHASE)
}
