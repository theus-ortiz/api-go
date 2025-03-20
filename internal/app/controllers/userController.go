package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/theus-ortiz/api-go/internal/app/config/logger"
	"github.com/theus-ortiz/api-go/internal/app/config/validation"
	"github.com/theus-ortiz/api-go/internal/app/models/requests"
	"github.com/theus-ortiz/api-go/internal/app/services"
	"go.uber.org/zap"
)

var validate = validator.New()

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := c.userService.GetAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var userReq requests.UserRequest

	if err := ctx.ShouldBindJSON(&userReq); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)
		ctx.JSON(errRest.Code, errRest)
		return
	}

	if err := validate.Struct(userReq); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validation.ValidateUserError(err)
		ctx.JSON(errRest.Code, errRest)
		return
	}

	if err := c.userService.CreateUser(userReq); err != nil {
		logger.Error("Error trying to create user", err,
			zap.String("journey", "createUser"),
		)
		ctx.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}