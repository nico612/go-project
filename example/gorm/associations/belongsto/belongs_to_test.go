package belongsto

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func initGorm() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:12345678@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
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
		&Company{},
	)
}

func TestFindUserByID(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)

	if err = repo.InsertCompany(&Company{Name: "company1", Code: "code1"}); err != nil {
		t.Fatal(err)
	}

	if err = repo.InsertUser(&User{Name: "user1", CompanyID: 1}); err != nil {
		t.Fatal(err)
	}

	// 1. 使用 Preload 方法加载关联数据
	user, err := repo.FindUserByIDWithPreload(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)

	// 2. 使用 Joins 方法手动连接表
	user, err = repo.FindUserByIDWithJoins(1)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("user: %+v", user)

	// 3. 使用 Select 方法选择字段
	username, err := repo.FindUserByIDWithSelectName(1)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("user: %s", username)
}
