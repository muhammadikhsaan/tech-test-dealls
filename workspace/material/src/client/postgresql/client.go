package postgresql

import (
	"context"
	"fmt"
	"net/http"

	"dealls.test/material/src/core"
	"dealls.test/material/src/static"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	MODE = static.MODE
)

type Fcx func(tx context.Context) *core.Error

type Connection struct {
	Conn *gorm.DB
	context.Context
}

type Client interface {
	Cnx(ctx context.Context) context.Context
	Trx(ctx context.Context, fc Fcx) *core.Error
}

type client struct {
	dbi *gorm.DB
}

var (
	host     = static.DATABASE_HOST
	port     = static.DATABASE_PORT
	user     = static.DATABASE_USER
	password = static.DATABASE_PASSWORD
	dbname   = static.DATABASE_NAME
)

func NewClient() (Client, error) {
	dbi, err := connection()

	if err != nil {
		return nil, err
	}

	return &client{
		dbi: dbi,
	}, nil
}

func connection() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), nil)

	if err != nil {
		return nil, err
	}

	if MODE == "DEBUG" {
		db = db.Debug()
	}

	return db, nil
}

func (p *client) connection() *core.Error {
	if p.dbi != nil {
		return nil
	}

	dbi, err := connection()

	if err != nil {
		return &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "failed to get db connection",
		}
	}

	p.dbi = dbi
	return nil
}

func (p *client) Trx(ctx context.Context, fc Fcx) *core.Error {

	if err := p.connection(); err != nil {
		return err
	}

	tx := p.dbi.Begin()

	err := fc(&Connection{
		Conn:    tx,
		Context: ctx,
	})

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (p *client) Cnx(ctx context.Context) context.Context {
	if err := p.connection(); err != nil {
		panic(err)
	}

	return &Connection{
		Conn:    p.dbi,
		Context: ctx,
	}
}
