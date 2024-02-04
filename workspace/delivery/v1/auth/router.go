package auth

import (
	"dealls.test/domain/src/usecase"
	"dealls.test/material/src/core"
	"dealls.test/material/src/middleware"
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
	h.routerGuest(r)

	r.Group(func(r core.Router) {
		r.Use(middleware.Accessable)
		h.routerAccess(r)
	})
}

func (h *handler) routerGuest(r core.Router) {
	r.Get("/", h.LOGIN)
	r.Post("/", h.REGISTER)

	r.Get("/check/email/{email}", h.CHECKEMAIL)
	r.Get("/check/username/{username}", h.CHECKUSERNAME)
}

func (h *handler) routerAccess(r core.Router) {
	r.Get("/me", h.ME)
	r.Delete("/", h.LOGOUT)
}
