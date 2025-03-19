package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/theus-ortiz/api-go/config/logger"
	"github.com/theus-ortiz/api-go/config/validation"
	"github.com/theus-ortiz/api-go/models"
	"github.com/theus-ortiz/api-go/models/requests"
	"go.uber.org/zap"
)

var validate = validator.New()

func GetAllUsers(c *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
    var userReq requests.UserRequest

    // Faz o bind dos dados da requisição para a struct UserRequest
    if err := c.ShouldBindJSON(&userReq); err != nil {
        logger.Error("Error trying to validate user info", err,
            zap.String("journey", "createUser"),
        )
        errRest := validation.ValidateUserError(err)
        c.JSON(errRest.Code, errRest)
        return
    }

    // Valida os dados da requisição
    if err := validate.Struct(userReq); err != nil {
        logger.Error("Error trying to validate user info", err,
            zap.String("journey", "createUser"),
        )
        errRest := validation.ValidateUserError(err)
        c.JSON(errRest.Code, errRest)
        return
    }

    // Cria o usuário no banco de dados
    user := models.User{}
    if err := user.CreateUser(userReq); err != nil {
        logger.Error(
            "Error trying to create user in database",
            err,
            zap.String("journey", "createUser"),
        )
        c.JSON(err.Code, err)
        return
    }

    logger.Info("User created successfully",
        zap.String("journey", "createUser"))

    // Retorna uma mensagem de sucesso
    c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}