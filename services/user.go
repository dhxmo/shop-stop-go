package service

import (
	"log"
	"net/http"

	"github.com/dhxmo/shop-stop-go/models"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/dhxmo/shop-stop-go/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

type UserService interface {
	GetUserByID(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type UserSvc struct {
	repo repositories.UserRepository
}

func NewUserSvc() UserService {
	return &UserSvc{repo: repositories.NewUserRepository()}
}

func (us *UserSvc) validate(r models.RegisterRequest) bool {
	return utils.Validate(
		[]utils.Validation{
			{Value: r.Username, Valid: "username"},
			{Value: r.Email, Valid: "email"},
			{Value: r.Password, Valid: "password"},
		})
}

func (us *UserSvc) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := us.repo.LoginUser(&req)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res models.UserResponse
	copier.Copy(&res, &user)
	res.Token = map[string]interface{}{
		"token": utils.GenerateToken(user),
	}
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

func (us *UserSvc) Register(c *gin.Context) {
	var reqBody models.RegisterRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := us.validate(reqBody)
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Request body is invalid"})
		return
	}

	user, err := us.repo.RegisterUser(&reqBody)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), ""))
		return
	}

	var res models.UserResponse
	copier.Copy(&res, &user)
	res.Token = map[string]interface{}{
		"token": utils.GenerateToken(user),
	}
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}

func (us *UserSvc) GetUserByID(c *gin.Context) {
	userUUID := c.Param("uuid")
	user, err := us.repo.GetUserByID(userUUID)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(http.StatusBadRequest, utils.Response(nil, err.Error(), "20001"))
		return
	}

	var res models.UserResponse
	copier.Copy(&res, &user)
	c.JSON(http.StatusOK, utils.Response(res, "OK", ""))
}
