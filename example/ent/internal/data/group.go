package data

import (
	"context"
	"github.com/nico612/go-project/example/ent/internal/data/ent/group"
	"log"
)

type GroupRepo interface {
	CreateGroupWithUser(ctx context.Context, groupName string, userId ...int) error
	QueryByName(ctx context.Context, groupName string) error
	QueryGroupWithUsers(ctx context.Context) error
}

type groupRepo struct {
	data *Data
}

func (g *groupRepo) CreateGroupWithUser(ctx context.Context, groupName string, userId ...int) error {
	return g.data.db.Group.Create().SetName(groupName).AddUserIDs(userId...).Exec(ctx)
}

func (g *groupRepo) QueryGroupWithUsers(ctx context.Context) error {
	groups, err := g.data.db.Group.Query().Where(group.HasUsers()).All(ctx)
	if err != nil {
		return err
	}

	log.Println("groups returned:", groups)
	return nil
}

func (g *groupRepo) QueryByName(ctx context.Context, groupName string) error {
	cars, err := g.data.db.Group.Query().Where(group.Name(groupName)).QueryUsers().QueryCars().All(ctx)
	if err != nil {
		return err
	}
	log.Println("cars returned: ", cars)
	return nil
}

func NewGroupRepo(data *Data) GroupRepo {
	return &groupRepo{data: data}
}
