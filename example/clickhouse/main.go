package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/nico612/go-project/example/clickhouse/db"
	"time"
)

type User struct {
	ID        int64     `ch:"id"`
	Username  string    `ch:"username"`
	Password  string    `ch:"password"`
	CreatedAt time.Time `ch:"created_at"`
	UpdatedAt time.Time `ch:"updated_at"`
	Period    uint32    `ch:"period"`
}

func main() {
	conn, err := db.GetClickHouseConn(&db.Options{
		ServerAddr:  "127.0.0.1:9000",
		Username:    "default",
		Password:    "",
		DBName:      "test",
		DialTimeOut: 10 * time.Second,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	fmt.Println("clickhouse connection success")

	ctx := context.Background()
	//user := &User{
	//	ID:        1,
	//	Username:  "zhangsan",
	//	Password:  "zhangsan123",
	//	CreatedAt: time.Now(),
	//	UpdatedAt: time.Now(),
	//	Period:    1,
	//}
	//
	//if err = insert(conn, ctx, user); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	users, err := queryUser(conn, ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(users)

	user, err := queryUserWithUserName(conn, ctx, "zhangsan")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("not found")
		} else {
			fmt.Println(err)
			return
		}

	}

	fmt.Println(user)
}

func insert(conn clickhouse.Conn, ctx context.Context, user *User) error {
	err := conn.Exec(ctx, `
		INSERT INTO test.user (id, username, password, created_at, updated_at, period)
		VALUES (?, ?, ?, ?, ?, ?)
	`, user.ID, user.Username, user.Password, user.CreatedAt, user.UpdatedAt, user.Period)
	return err
}

func queryUser(conn clickhouse.Conn, ctx context.Context) ([]User, error) {
	var users []User
	err := conn.Select(ctx, &users, "SELECT * FROM user")
	return users, err
}

func queryUserWithUserName(conn clickhouse.Conn, ctx context.Context, username string) (*User, error) {
	var user User
	err := conn.QueryRow(ctx, "SELECT * FROM test.user WHERE username = ?", username).ScanStruct(&user)

	return &user, err

}
