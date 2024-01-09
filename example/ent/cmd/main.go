package main

import (
	"context"
	"fmt"
	"github.com/nico612/go-project/example/ent/internal/data"
	"log"
)

func main() {
	client := data.NewEntClient("mysql", "root:123456@tcp(127.0.0.1)/ent-test?parseTime=true")
	d := data.NewData(client)

	QueryGroupWithUsers(d)
}

func queryCarOwner(d *data.Data) {
	carRepo := data.NewCarsRepo(d)
	if err := carRepo.QueryCarUser(context.Background(), "lisi"); err != nil {
		log.Fatal(err)
	}
}

func UserOperate(d *data.Data) {
	userRepo := data.NewUserRepo(d)

	userNmae := "lisi"

	// 1. 创建用户
	u, err := userRepo.CreateUser(context.Background(), userNmae, 18)
	if err != nil {
		panic(err)
	}

	fmt.Printf("create user = %v", u)

	// 2. 查询用户
	u, err = userRepo.QueryUser(context.Background(), userNmae)
	if err != nil {
		panic(err)
	}

	fmt.Printf("query user = %v \n", u)

	// 3. 为用户添加汽车
	userM, err := userRepo.UpdateUserCars(context.Background(), "byd", userNmae)
	if err != nil {
		panic(err)
	}
	fmt.Printf("update user = %v \n", userM)
}

func TestGroupManyToMany(d *data.Data) {

	// 1. 创建user
	userRepo := data.NewUserRepo(d)
	ariel, err := userRepo.CreateUser(context.Background(), "ariel", 30)
	if err != nil {
		panic(err)
	}

	neta, err := userRepo.CreateUser(context.Background(), "Neta", 28)
	if err != nil {
		panic(err)
	}

	// 2. 为用户创建car
	carRepo := data.NewCarsRepo(d)
	if err = carRepo.CreateCarToUser(context.Background(), "Tesla", ariel.ID); err != nil {
		panic(err)
	}

	if err = carRepo.CreateCarToUser(context.Background(), "Mazda", ariel.ID); err != nil {
		panic(err)
	}

	if err = carRepo.CreateCarToUser(context.Background(), "Ford", neta.ID); err != nil {
		panic(err)
	}

	//3 创建 group 并将用户加入 group
	groupRepo := data.NewGroupRepo(d)
	if err = groupRepo.CreateGroupWithUser(context.Background(), "Gitlab", ariel.ID, neta.ID); err != nil {
		panic(err)
	}

	if err = groupRepo.CreateGroupWithUser(context.Background(), "Github", ariel.ID); err != nil {
		panic(err)
	}

	log.Println("The graph was created successfully")

}

func QueryGithub(d *data.Data) {
	grouprepo := data.NewGroupRepo(d)
	if err := grouprepo.QueryByName(context.Background(), "Github"); err != nil {
		panic(err)
	}
}

func QueryGroupWithUsers(d *data.Data) {
	grouprepo := data.NewGroupRepo(d)
	if err := grouprepo.QueryGroupWithUsers(context.Background()); err != nil {
		panic(err)
	}
}
