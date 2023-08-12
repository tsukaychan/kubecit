package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"kubecit/ent"
	"kubecit/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	conf *conf.Data
	db   *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	entClient, err := ent.Open(c.Database.Driver, c.Database.Source)
	if err != nil {
		log.Fatalf("fail to open connection to db,%s", err)
	}
	if err = entClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("fail to create schema,%s", err)
	}
	return &Data{
		conf: c,
		db:   entClient,
	}, cleanup, nil
}
