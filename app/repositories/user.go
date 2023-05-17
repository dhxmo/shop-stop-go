package repositories

import (
	"errors"

	"github.com/dhxmo/shop-stop-go/app/models"
	"github.com/dhxmo/shop-stop-go/config"
	"github.com/dhxmo/shop-stop-go/pkg/utils"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetUserByID(uuid string) (*models.UserResponse, error)
	RegisterUser(req *models.RegisterRequest) (*models.UserResponse, error)
	LoginUser(req *models.LoginRequest) (*models.UserResponse, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepo{db: config.DB}
}

func (ur *UserRepo) GetUserByID(uuid string) (*models.UserResponse, error) {
	var user models.User
	if err := ur.db.Where("uuid = ?", uuid).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.UUID == "" {
		return nil, nil
	}

	var res models.UserResponse
	copier.Copy(&res, &user)

	return &res, nil
}

func (ur *UserRepo) RegisterUser(req *models.RegisterRequest) (*models.UserResponse, error) {
	var user models.User
	copier.Copy(&user, &req)

	hashedPassword := utils.HashAndSalt([]byte(req.Password))
	user.Password = hashedPassword

	if err := ur.db.Create(&user).Error; err != nil {
		return nil, err
	}

	var res models.UserResponse
	copier.Copy(&res, &user)

	return &res, nil
}

func (ur *UserRepo) LoginUser(req *models.LoginRequest) (*models.UserResponse, error) {
	user := &models.User{}
	if err := ur.db.Where("username = ? ", req.Username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return nil, errors.New("wrong password")
	}

	var res models.UserResponse
	copier.Copy(&res, &user)

	return &res, nil
}
