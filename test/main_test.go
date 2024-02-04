package main_test

import (
	"context"
	"testing"

	"dealls.test/delivery/v1/auth"
	"dealls.test/delivery/v1/interaction"
	"dealls.test/delivery/v1/purchase"
	"dealls.test/domain/src/data"
	"dealls.test/domain/src/usecase"
	"dealls.test/material/src/client"
	"dealls.test/material/src/core"
	"dealls.test/material/src/secret"
)

const (
	authRoute        = "/testing/auth"
	purchaseRoute    = "/testing/purchase"
	interactionRoute = "/testing/interaction"
)

var (
	router *core.Mux
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	secret.LoadSecretKeyJWT()

	router = core.NewRouter()

	c, _ := client.NewClient(ctx)

	dt := data.NewRepository()
	uc := usecase.NewService(dt, c)

	router.Route(authRoute, auth.NewHandler(uc).Router)
	router.Route(purchaseRoute, purchase.NewHandler(uc).Router)
	router.Route(interactionRoute, interaction.NewHandler(uc).Router)
	m.Run()
}
