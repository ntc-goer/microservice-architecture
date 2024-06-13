package pkg

import (
	"context"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/ent"
)

type DB struct {
	Config *config.Config
	Client *ent.Client
}

func NewDB(cfg *config.Config) (*DB, error) {
	client, err := ent.Open("postgres", "host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		return nil, err
	}
	return &DB{
		Config: cfg,
		Client: client,
	}, nil
}

func (db *DB) MigrateDatabase() error {
	err := db.Client.Schema.Create(context.Background())
	return err
}
