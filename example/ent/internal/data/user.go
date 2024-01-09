package data

import (
	"context"
	"github.com/nico612/go-project/example/ent/internal/data/ent"
	"github.com/nico612/go-project/example/ent/internal/data/ent/user"
	"time"
)

type UserRepo interface {
	CreateUser(ctx context.Context, name string, age int) (*ent.User, error)
	QueryUser(ctx context.Context, name string) (*ent.User, error)
	UpdateUserCars(ctx context.Context, carModel string, userName string) (*ent.User, error)
}

type userRepo struct {
	data *Data
}

func NewUserRepo(data *Data) UserRepo {
	return &userRepo{data: data}
}

func (r *userRepo) CreateUser(ctx context.Context, name string, age int) (*ent.User, error) {
	return r.data.db.User.Create().SetAge(age).SetName(name).Save(ctx)
}

func (r *userRepo) QueryUser(ctx context.Context, name string) (*ent.User, error) {
	u, err := r.data.db.User.Query().Where(user.Name(name)).Only(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// UpdateUserCars 为用户添加Cars
func (r *userRepo) UpdateUserCars(ctx context.Context, carModel string, userName string) (*ent.User, error) {
	// 1. 查询user是否存在
	u, err := r.QueryUser(ctx, userName)
	if err != nil {
		return nil, err
	}

	//2. 创建一个car
	car, err := r.data.db.Car.
		Create().
		SetModel(carModel).
		SetRegisteredAt(time.Now().String()).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// 更新user
	u, err = r.data.db.User.UpdateOne(u).AddCars(car).Save(ctx)
	if err != nil {
		return nil, err
	}

	return u, nil
}
