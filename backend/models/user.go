package models

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"-" gorm:"not null"` // "-" 表示不在 JSON 中显示
	Todos    []Todo `json:"todos" gorm:"foreignKey:UserID"`
}

// 更新 Todo 模型以包含用户关联
func UpdateTodoModel() {
	// 在 models/todo.go 中添加 UserID 字段
}

// HashPassword 加密用户密码
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword 验证用户密码
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}