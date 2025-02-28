package users

import (
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) *repository {
	return &repository{
		db: db,
	}
}

func (repo *repository) All() ([]User, error) {
	var users []User
	err := repo.db.Find(&users).Error
	return users, err
}

func (repo *repository) Store(user *User) error {
	return repo.db.Create(&user).Error
}

func (repo *repository) Show(id uint) (*User, error) {
	var user User
	err := repo.db.First(&user, id).Error
	return &user, err
}

func (repo *repository) Update(id uint, user *User) error {
	return repo.db.
		Model(&User{ID: id}).
		Updates(user).Error
}

func (repo *repository) Destroy(id uint) error {
	return repo.db.Delete(&User{}, id).Error
}
