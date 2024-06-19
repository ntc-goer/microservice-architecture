package repository

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ntc-goer/microservice-examples/orderservice/config"
	"github.com/ntc-goer/microservice-examples/orderservice/ent"
)

type Repository struct {
	Client *ent.Client
	Order  *OrderRepo
	Dish   *DishRepo
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.ServerHost,
			cfg.Database.ServerPort,
			cfg.Database.UserName,
			cfg.Database.Password,
			cfg.Database.DBName,
		))
	if err != nil {
		return nil, err
	}
	return &Repository{
		Client: client,
		Order:  NewOrderRepo(client.Order),
		Dish:   NewDishRepo(client.Dish),
	}, nil
}

func (r *Repository) MigrateDatabase() error {
	err := r.Client.Schema.Create(context.Background())
	return err
}
