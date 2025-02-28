package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewController(srvs IService) *controller {
	return &controller{
		srvs: srvs,
	}
}

func (ctrl *controller) Index(ctx *gin.Context) {
	users, err := ctrl.srvs.All()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users})
}

func (ctrl *controller) Store(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if err := ctrl.srvs.Store(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl *controller) Show(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	user, err := ctrl.srvs.Show(uint(id))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": "User not found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (ctrl *controller) Update(ctx *gin.Context) {
	var user User
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if err := ctrl.srvs.Update(uint(id), &user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (ctrl *controller) Destroy(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctrl.srvs.Destroy(uint(id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}
