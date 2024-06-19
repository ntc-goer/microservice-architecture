package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/ntc-goer/microservice-examples/orderservice/ent"
)

type DishRepo struct {
	DishClient *ent.DishClient
}

type DishItem struct {
	DishName string
	DishId   string
	Quantity int
}

func NewDishRepo(dc *ent.DishClient) *DishRepo {
	return &DishRepo{
		DishClient: dc,
	}
}

func (dr *DishRepo) CreateDishes(ctx context.Context, orderId uuid.UUID, dishes []*DishItem) error {
	_, err := dr.DishClient.MapCreateBulk(dishes, func(c *ent.DishCreate, i int) {
		c.SetOrderID(orderId).SetDishID(dishes[i].DishId).SetDishName(dishes[i].DishName).SetQuantity(dishes[i].Quantity)
	}).Save(ctx)
	return err
}
