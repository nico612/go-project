package belongsto

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string
	Code string
}

// TableName 重写表名
func (Company) TableName() string {
	return "companies"
}

// User 属于 Company, CompanyID 是外键
// 默认情况下， CompanyID 被隐含地用来在 User 和 Company 之间创建一个外键关系，
// 因此必须包含在 User 结构体中才能填充 Company 内部结构体。
// 要定义一个 belongs to 关系，数据库的表中必须存在外键。默认情况下，外键的名字，使用拥有者的类型名称加上表的主键的字段名字
// 重写外键 `gorm:"foreignKey:CompanyID"`
// 如果设置了User实体属于Company实体，那么GORM会自动把Company中的ID属性保存到User的CompanyID属性中。
// 重写引用 `gorm:"references:Code"` 使用Code作为引用
// NOTE: 如果外键名恰好在拥有者类型中存在，GORM 通常会错误的认为它是 has one 关系。我们需要在 belongs to 关系中指定 references
// 你可以通过 constraint 标签配置 OnUpdate、OnDelete 实现外键约束，在使用 GORM 进行迁移时它会被创建，例如：
// Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
type User struct {
	gorm.Model
	Name      string
	CompanyID uint
	Company   Company `gorm:"foreignKey:CompanyID"`
}

func (User) TableName() string {
	return "users"
}

// 关联查询

type Repo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}

// insert Company
func (r *Repo) InsertCompany(company *Company) error {
	return r.db.Create(company).Error
}

// delete Company
func (r *Repo) DeleteCompany(id int) error {
	return r.db.Delete(&Company{}, id).Error
}

// insert User
func (r *Repo) InsertUser(user *User) error {
	return r.db.Create(user).Error
}

// 使用 Preload 方法加载关联数据
func (r *Repo) FindUserByIDWithPreload(id int) (*User, error) {
	var user User
	if err := r.db.Preload("Company").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 使用 Joins 方法手动连接表
func (r *Repo) FindUserByIDWithJoins(id int) (*User, error) {
	var user User
	if err := r.db.Joins("Company").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// 3. 使用 Select 方法选择字段
func (r *Repo) FindUserByIDWithSelectName(id int) (string, error) {
	var username string
	if err := r.db.Model(&User{}).Select("name").First(&username, id).Error; err != nil {
		return "nil", err
	}
	return username, nil
}

// 4. 使用 Scopes 方法定义查询条件
// 5. 使用 Association 方法加载关联数据
// 6. 使用 Related 方法加载关联数据
// 7. 使用 Count 方法统计关联数据
// 8. 使用 Find 方法查询关联数据
// 9. 使用 Append 方法追加关联数据
// 10. 使用 Replace 方法替换关联数据
// 11. 使用 Delete 方法删除关联数据
// 12. 使用 Clear 方法清除关联数据
// 13. 使用 DBClause 方法自定义关联查询
// 14. 使用 Set 方法设置关联数据
// 15. 使用 Add 方法添加关联数据
// 16. 使用 Remove 方法移除关联数据
// 17. 使用 Replace 方法替换关联数据
// 18. 使用 FindInBatches 方法批量查询关联数据
// 19. 使用 CountInBatches 方法批量统计关联数据
// 20. 使用 Preload 方法预加载关联数据
// 21. 使用 Preload 方法预加载多个关联数据
