package controllers

import (
	"net/http"

	models "github.com/dhxmo/shop-stop-go/app/models"
	services "github.com/dhxmo/shop-stop-go/app/services"
	utils "github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type User struct {
	service services.UserService
}

func NewUserController(service services.UserService) *User {
	return &User{
		service: service,
	}
}

func (us *User) validate(r models.RegisterRequest) (bool, error) {
	val := utils.Validate(
		[]utils.Validation{
			{Value: r.Username, Valid: "username"},
			{Value: r.Email, Valid: "email"},
			{Value: r.Password, Valid: "password"},
		})

	return val, nil
}

func (us *User) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	user, err := us.service.Login(ctx, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res models.UserResponse
	copier.Copy(&res, &user)

	token, err := utils.GenerateToken(user)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

func (us *User) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := us.validate(req)
	if !valid || err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is invalid"})
		return
	}

	ctx := c.Request.Context()
	user, err := us.service.Register(ctx, &req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, "", err.Error()))
		return
	}

	var res models.UserResponse
	copier.Copy(&res, &user)

	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

func (us *User) GetUserByID(c *gin.Context) {
	userUUID := c.Param("uuid")
	ctx := c.Request.Context()

	user, err := us.service.GetUserByID(ctx, userUUID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), "20001"))
		return
	}

	var res models.UserResponse
	copier.Copy(&res, &user)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}
