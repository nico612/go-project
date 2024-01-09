package data

import (
	"context"
	"github.com/nico612/go-project/example/ent/internal/data/ent/user"
	"log"
	"time"
)

type CarsRepo interface {
	QueryCarUser(ctx context.Context, userName string) error
	CreateCarToUser(ctx context.Context, model string, userId int) error
}

type carsRepo struct {
	data *Data
}

func (c *carsRepo) CreateCarToUser(ctx context.Context, model string, userId int) error {
	return c.data.db.Car.Create().
		SetModel(model).
		SetRegisteredAt(time.Now().String()).
		SetOwnerID(userId).
		Exec(ctx)
}

func (c *carsRepo) QueryCarUser(ctx context.Context, userName string) error {
	u, err := c.data.db.User.Query().Where(user.Name(userName)).Only(ctx)
	if err != nil {
		return err
	}

	cars, err := c.data.db.User.QueryCars(u).All(ctx)
	if err != nil {
		return err
	}

	for _, car := range cars {
		owner, err := car.QueryOwner().Only(ctx)
		if err != nil {
			return err
		}

		log.Printf("car %q owner: %q\n", car.Model, owner.Name)
	}
	return nil
}

func NewCarsRepo(data *Data) CarsRepo {
	return &carsRepo{data: data}
}
