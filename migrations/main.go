package main

import (
	"context"
	"fmt"

	"dealls.test/domain/src/data/interactions"
	"dealls.test/domain/src/data/privilages"
	"dealls.test/domain/src/data/users"
	"dealls.test/material/src/client/postgresql"
)

func main() {
	ctx := context.Background()
	c, _ := postgresql.NewClient()

	dbi := c.Cnx(ctx).(*postgresql.Connection).Conn

	if err := users.Migrate(dbi).Executor(); err != nil {
		fmt.Println(err)
	}

	if err := privilages.Migrate(dbi).Executor(); err != nil {
		fmt.Println(err)
	}

	if err := interactions.Migrate(dbi).Executor(); err != nil {
		fmt.Println(err)
	}
}
