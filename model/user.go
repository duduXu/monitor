package model

type User struct {
	ID       uint64 `gorm:"primary_key"`
	Username string
	Password string
}

func (User) TableName() string {
	return "user"
}
