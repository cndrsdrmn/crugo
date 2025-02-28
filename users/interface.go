package users

import "github.com/gin-gonic/gin"

type IController interface {
	Index(ctx *gin.Context)
	Store(ctx *gin.Context)
	Show(ctx *gin.Context)
	Update(ctx *gin.Context)
	Destroy(ctx *gin.Context)
}

type IRepository interface {
	All() ([]User, error)
	Store(user *User) error
	Show(id uint) (*User, error)
	Update(id uint, user *User) error
	Destroy(id uint) error
}

type IService interface {
	All() ([]User, error)
	Store(user *User) error
	Show(id uint) (*User, error)
	Update(id uint, user *User) error
	Destroy(id uint) error
}
