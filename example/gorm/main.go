package main

import (
	"github.com/nico612/go-project/example/gorm/associations/belongsto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initGorm()

}

// init gorm
func initGorm() *gorm.DB {
	dns := "root:12345678@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时删除旧索引，然后创建一个新索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err = autoMigrate(db); err != nil {
		panic(err)
	}
	return db
}

func autoMigrate(db *gorm.DB) error {
	// 自动迁移
	return db.AutoMigrate(
		&belongsto.User{},
		&belongsto.Company{},
	)
}
