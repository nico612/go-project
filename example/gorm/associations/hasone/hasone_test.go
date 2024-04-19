package hasone

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

// hash one 基本操作
func TestHasOne(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)

	// insert User
	if err = repo.InsertUser(&User{Name: "user1"}); err != nil {
		t.Fatal(err)
	}

	// insert CreditCard
	if err = repo.InsertCreditCard(&CreditCard{Number: "12345678", UserID: 1}); err != nil {
		t.Fatal(err)
	}

	// find user by id with preload， 预加载 CreditCard
	user, err := repo.FindUserByIDWithPreload(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)

	// find users with preload
	users, err := repo.FindUsersWithPreload()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("users: %+v", users)
}

// join 查询
func TestHasOneWithJoins(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)

	users, err := repo.FindUsersWithJoins()
	if err != nil {
		t.Fatal(err)
	}
	for i, user := range users {
		t.Logf("user[%d]: %+v", i, user)
	}
}

// 多态
func TestPolymorphic(t *testing.T) {

	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)

	dog := &Dog{Name: "dog2", Toy: Toy{Name: "toy1"}}
	cat := &Cat{Name: "cat2", Toy: Toy{Name: "toy2"}}
	err = repo.InsertDog(dog)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.InsertCat(cat)
	if err != nil {
		t.Fatal(err)
	}
}

// 多态关联查询
func TestPolymorphicWithPreload(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)

	cat, err := repo.FindCatByIDWithPreload(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("cat: %+v", cat)
}

// 多态更新
func TestPolymorphicUpdate(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)

	cat, err := repo.FindCatByIDWithPreload(1)
	if err != nil {
		t.Fatal(err)
	}

	cat.Toy.Name = "toy3"
	err = repo.UpdateToy(&cat.Toy)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cat: %+v", cat)

	cat, err = repo.FindCatByIDWithPreload(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("cat: %+v", cat)

}

func initGorm() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:12345678@tcp(127.0.0.1:3306)/gorm_hasone?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时删除旧索引，然后创建一个新索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = autoMigrate(db); err != nil {
		return nil, err
	}
	return db, err
}

func autoMigrate(db *gorm.DB) error {
	// 自动迁移
	return db.AutoMigrate(
		&User{},
		&CreditCard{},
		&Cat{},
		&Dog{},
		&Toy{},
	)
}
