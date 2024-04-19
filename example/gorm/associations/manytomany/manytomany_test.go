package manytomany

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestManyToMany(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)

	// insert User
	user := &User{
		Name: "user1",
		Languages: []Language{
			{Name: "english"},
			{Name: "chinese"},
			{Name: "french"},
		},
	}
	if err = repo.InsertUser(user); err != nil {
		t.Fatal(err)
	}

	t.Logf("user: %+v\n", user)

	// find user by id with preload
	user, err = repo.FindUserByIDWithPreload(int(user.ID))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("user: %+v\n", user)

	// find users with preload
	users, err := repo.FindUsersWithPreload()
	if err != nil {
		t.Fatal(err)
	}

	for i, user := range users {
		t.Logf("user[%d]: %+v\n", i, user)
	}

}

func TestManyToMany2(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}
	repo := NewUserRepo(db)
	// insert User2
	user := &User{
		Name: "user2",
	}
	if err = repo.InsertUser(user); err != nil {
		t.Fatal(err)
	}

	t.Logf("user: %+v\n", user)

	// insert Language
	// update user2 languages
	if err = repo.InsertUserLanguages(&User{
		Model: gorm.Model{ID: 2},
	}, []*Language{
		{Name: "japanese"},
	}); err != nil {
		t.Fatal(err)
	}

	// find user by id with preload
	user, err = repo.FindUserByIDWithPreload(int(user.ID))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("user: %+v\n", user)

	// find users with preload
	users, err := repo.FindUsersWithPreload()
	if err != nil {
		t.Fatal(err)
	}

	for i, user := range users {
		t.Logf("user[%d]: %+v\n", i, user)
	}
}

func TestManyToMany3(t *testing.T) {
	// 为用户添加语言
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)
	}

	repo := NewUserRepo(db)

	user, err := repo.FindUserByIDWithPreload(2)
	if err != nil {
		t.Fatal(err)
	}

	// 添加语言
	languages := []*Language{
		{
			Model: gorm.Model{ID: 3},
		},
	}

	if err = repo.InsertUserLanguages(user, languages); err != nil {
		t.Fatal(err)
	}

}

// test find languages
func TestFindLanguages(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)

	}
	repo := NewUserRepo(db)

	languages, err := repo.FindLanguagesWithPreload()
	if err != nil {
		t.Fatal(err)

	}
	for _, language := range languages {
		t.Logf("language: %+v\n", language)
	}
}

// delete user language
func TestDeleteUserLanguage(t *testing.T) {
	db, err := initGorm()
	if err != nil {
		t.Fatal(err)

	}
	repo := NewUserRepo(db)

	// 这种删除不会删除关联表中的数据
	err = repo.DeleteUser(2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteUserLanguage2(t *testing.T) {

	db, err := initGorm()
	if err != nil {
		t.Fatal(err)

	}
	repo := NewUserRepo(db)

	user, err := repo.FindUserByIDWithPreload(2)
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteUserWithLanguage(int(user.ID))
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteUserAndLanguage2(user)
	if err != nil {
		t.Fatal(err)
	}
}

func initGorm() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:12345678@tcp(127.0.0.1:3306)/gorm_mtm?charset=utf8mb4&parseTime=True&loc=Local",
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
		&Language{},
	)
}
