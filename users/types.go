package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type controller struct {
	srvs IService
}

type repository struct {
	db *gorm.DB
}

type service struct {
	repo IRepository
}

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"index:idx_name_users"`
	Email     string         `json:"email" gorm:"index:idx_email_users,unique"`
	Password  string         `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (m *User) BeforeCreate(tx *gorm.DB) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	m.Password = string(hashed)
	return nil
}
