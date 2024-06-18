package pkg

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/ent"
)

type DB struct {
	Config *config.Config
	Client *ent.Client
}

func NewDB(cfg *config.Config) (*DB, error) {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.DatabaseServerHost,
			cfg.DatabaseServerPort,
			cfg.DatabaseUser,
			cfg.DatabasePwd,
			cfg.DatabaseName,
		))
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
