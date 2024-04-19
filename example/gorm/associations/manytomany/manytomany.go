package manytomany

import "gorm.io/gorm"

// Many To Many
// Many to Many 会在两个 model 中添加一张连接表。
// 例如，您的应用包含了 user 和 language，且一个 user 可以说多种 language，多个 user 也可以说一种 language。
// 当使用 GORM 的 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表

// User 拥有并属于多个 Language，`user_languages` 是连接表
type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

// TableName 重写表名
func (User) TableName() string {
	return "users"
}

type Language struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_languages;"` // 反向引用
}

// TableName 重写表名
func (Language) TableName() string {
	return "languages"
}

type Repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

// InsertUser insert user
func (r *Repo) InsertUser(user *User) error {
	return r.db.Create(user).Error
}

// InsertLanguage insert language
func (r *Repo) InsertLanguage(language *Language) error {
	return r.db.Create(language).Error
}

// InsertUserLanguages insert user languages
func (r *Repo) InsertUserLanguages(user *User, languages []*Language) error {
	return r.db.Model(user).Association("Languages").Append(languages)
}

// FindUserByIDWithPreload find user by id with preload
func (r *Repo) FindUserByIDWithPreload(id int) (*User, error) {
	var user User
	if err := r.db.Preload("Languages").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUsersWithPreload find users with preload
func (r *Repo) FindUsersWithPreload() ([]*User, error) {
	var users []*User
	if err := r.db.Preload("Languages").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindUsersWithPreload find users with preload
func (r *Repo) FindUsersWithPreloadAndSelect() ([]*User, error) {
	var users []*User
	if err := r.db.Preload("Languages", "name = ?", "english").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// find languages with preload
func (r *Repo) FindLanguagesWithPreload() ([]*Language, error) {
	var languages []*Language
	if err := r.db.Preload("Users").Find(&languages).Error; err != nil {
		return nil, err
	}
	return languages, nil
}

// delete user
func (r *Repo) DeleteUser(id int) error {
	return r.db.Delete(&User{}, id).Error
}

// delete language when delete user
func (r *Repo) DeleteUserWithLanguage(id int) error {
	user := &User{Model: gorm.Model{ID: uint(id)}}
	return r.db.Select("Languages").Delete(user).Error
}

// delete user
func (r *Repo) DeleteUserAndLanguage2(user *User) error {
	return r.db.Select("Languages").Delete(user).Error
}
