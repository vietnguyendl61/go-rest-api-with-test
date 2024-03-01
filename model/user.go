package model

type User struct {
	BaseModel
	UserName string `json:"user_name" gorm:"column:user_name;type:varchar(255);not null"`
	Password string `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Address  string `json:"address" gorm:"column:address;type:text"`
	Email    string `json:"email" gorm:"column:email"`
}

func (u User) TableName() string {
	return "users"
}
