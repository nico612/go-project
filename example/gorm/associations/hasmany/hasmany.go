package hasmany

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard
}

// TableName rewrite table name
func (User) TableName() string {
	return "users"
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}

// table name
func (CreditCard) TableName() string {
	return "credit_cards"
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

// insert credit cards
func (r *Repo) InsertCreditCards(creditCards []*CreditCard) error {
	return r.db.Create(creditCards).Error
}

// FindUserByIDWithPreload find user by id with preload
func (r *Repo) FindUserByIDWithPreload(id int) (*User, error) {
	var user User
	if err := r.db.Preload("CreditCards").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindUsersWithPreload find users with preload
func (r *Repo) FindUsersWithPreload() ([]*User, error) {
	var users []*User
	if err := r.db.Preload("CreditCards").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// ********* 多态关联 参考官方文档，跟hasone 类似 **************/
