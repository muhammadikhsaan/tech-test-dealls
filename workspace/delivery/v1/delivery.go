package v1

import (
	"dealls.test/delivery/v1/auth"
	"dealls.test/delivery/v1/interaction"
	"dealls.test/delivery/v1/purchase"
	"dealls.test/domain/src/usecase"
	"dealls.test/material/src/core"
	"dealls.test/material/src/middleware"
)

type Delivery interface {
	Router(r core.Router)
}

type delivery struct {
	uc *usecase.Service
}

func NewDelivery(uc *usecase.Service) Delivery {
	return &delivery{
		uc: uc,
	}
}

func (c *delivery) Router(r core.Router) {
	// GUEST
	r.Group(func(r core.Router) {
		// r.Use(middleware.Throttle)
		r.Route("/auth", auth.NewHandler(c.uc).Router)
	})

	// USERS
	r.Group(func(r core.Router) {
		// r.Use(middleware.Throttle)
		r.Use(middleware.Accessable)
		r.Route("/interaction", interaction.NewHandler(c.uc).Router)
		r.Route("/purchase", purchase.NewHandler(c.uc).Router)
	})
}
