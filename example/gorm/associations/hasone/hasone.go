package hasone

import "gorm.io/gorm"

// has one 与另一个模型建立一对一的关联，但它和一对一关系有些许不同。
// 这种关联表明一个模型的每个实例都包含或拥有另一个模型的一个实例。

// User has one CreditCard, CreditCardID is the foreign key
// 重写外键 CreditCard `gorm:"foreignKey:UserName"` // 使用 UserName 作为外键
// 重写引用 CreditCard `gorm:"foreignKey:UserName;references:name"` // 使用 Name 作为引用
type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard // CreditCard 属于 User, UserID 是外键
}

// TableName rewrite table name
func (User) TableName() string {
	return "users"
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
	//UserName string
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

// InsertCreditCard insert credit card
func (r *Repo) InsertCreditCard(creditCard *CreditCard) error {
	return r.db.Create(creditCard).Error
}

// FindUserByIDWithPreload find user by id with preload
func (r *Repo) FindUserByIDWithPreload(id int) (*User, error) {
	var user User
	if err := r.db.Preload("CreditCard").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// users
func (r *Repo) FindUsersWithPreload() ([]*User, error) {
	var users []*User
	if err := r.db.Preload("CreditCard").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// find users with join
func (r *Repo) FindUsersWithJoins() ([]*User, error) {
	var users []*User
	// sql 语句 = SELECT * FROM `users` INNER JOIN `credit_cards` ON `credit_cards`.`user_id` = `users`.`id`
	if err := r.db.Joins("CreditCard").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

/****************** 多态关联 ********************/
// 多态关联是指 一个模型可以关联多个模型，这些模型可以是不同的模型，但是它们之间有一些共同的特性。
// 多态关联需要两个字段，一个是关联的模型的类型，另一个是关联的模型的 ID。例如下面 Toy 模型的 OwnerType 和 OwnerID 字段。
// OwnerType 用于存储关联的模型的类型，默认为关联的模型的表名，OwnerID 用于存储关联的模型的 ID，例如 Cat 和 Dog 的模型。

// 多态关联
type Cat struct {
	gorm.Model
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"` // 可以通过 `polymorphicValue:master` 来指定多态关联的值, 也就是 OwnerType 的值 为 master
}

// TableName rewrite table name
func (Cat) TableName() string {
	return "cats"
}

type Dog struct {
	gorm.Model
	Name string
	Toy  Toy `gorm:"polymorphic:Owner"`
}

// table name
func (Dog) TableName() string {
	return "dogs"
}

type Toy struct {
	gorm.Model
	Name      string
	OwnerID   uint
	OwnerType string
}

// table name
func (Toy) TableName() string {
	return "toys"
}

// InsertCat insert cat
func (r *Repo) InsertCat(cat *Cat) error {
	return r.db.Create(cat).Error
}

// InsertDog insert dog
func (r *Repo) InsertDog(dog *Dog) error {
	return r.db.Create(dog).Error
}

// find cat by id
func (r *Repo) FindCatByIDWithPreload(id int) (*Cat, error) {
	var cat Cat
	if err := r.db.Preload("Toy").First(&cat, id).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

// 多态关联更新, 只需要更新 toy 表中的字段即可
func (r *Repo) UpdateToy(toy *Toy) error {
	return r.db.Save(toy).Error
}
