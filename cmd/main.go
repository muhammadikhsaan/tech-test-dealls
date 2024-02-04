package main

import (
	"context"
	"fmt"
	"net/http"

	"dealls.test/domain/src/data"
	"dealls.test/domain/src/usecase"
	"dealls.test/material/src/client"
	"dealls.test/material/src/core"
	pm "dealls.test/material/src/middleware"
	"dealls.test/material/src/secret"
	"dealls.test/material/src/static"
	"github.com/go-chi/chi/v5/middleware"

	v1 "dealls.test/delivery/v1"
)

var (
	MODE = static.MODE
	PORT = fmt.Sprintf(":%s", static.PORT)
)

func main() {
	secret.LoadSecretKeyJWT()

	r := core.NewRouter()
	ctx := context.Background()

	if MODE == "DEBUG" {
		r.Use(middleware.Logger)
	}

	r.Use(pm.Recovery)
	r.Use(middleware.CleanPath)
	r.Use(pm.Cors)

	c, status := client.NewClient(ctx)
	dt := data.NewRepository()
	uc := usecase.NewService(dt, c)

	r.Route("/api/v1", v1.NewDelivery(uc).Router)

	fmt.Printf("Client Status : %v \n", status)
	fmt.Printf("Server running on port%s \n", PORT)
	if err := http.ListenAndServe(PORT, r); err != nil {
		ctx.Done()
		fmt.Println(err.Error())
	}
}
