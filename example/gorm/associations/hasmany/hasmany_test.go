package hasmany

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestHasMany(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)

	// 插入用户
	user := &User{
		Name: "zhangsan",
	}

	if err = repo.InsertUser(user); err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)

	user2 := &User{
		Name: "lisi",
	}
	if err = repo.InsertUser(user2); err != nil {
		t.Fatal(err)
	}

	// 插入信用卡
	creditCards := []*CreditCard{
		{
			Number: "12345678",
			UserID: user.ID,
		},
		{
			Number: "87654321",
			UserID: user.ID,
		},
		{
			Number: "11111111",
			UserID: user2.ID,
		},
		{
			Number: "22222222",
			UserID: user2.ID,
		},
	}

	if err = repo.InsertCreditCards(creditCards); err != nil {
		t.Fatal(err)
	}

	// 通过预加载 CreditCards 查找用户
	user, err = repo.FindUserByIDWithPreload(int(user.ID))
	if err != nil {
		t.Fatal(err)
	}

	// 通过预加载 CreditCards 查找所有用户
	users, err := repo.FindUsersWithPreload()
	if err != nil {
		t.Fatal(err)
	}

	for i, user := range users {
		t.Logf("user[%d]: %+v", i, user)
	}

}

func initGorm() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:12345678@tcp(127.0.0.1:3306)/gorm_hasmany?charset=utf8mb4&parseTime=True&loc=Local",
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
	)
}
