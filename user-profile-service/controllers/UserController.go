package controllers

import (
	"net/http"
	"strconv"

	models "github.com/Aranyak-Ghosh/spotigo/user_profile/models/database/user"
	httpModels "github.com/Aranyak-Ghosh/spotigo/user_profile/models/http"
	service "github.com/Aranyak-Ghosh/spotigo/user_profile/services/user"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type userController struct {
	service service.UserService
	logger  *zap.SugaredLogger
}

type UserController interface {
	ListUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	FollowUser(c *gin.Context)
	UnFollowUser(c *gin.Context)
	ListFollowers(c *gin.Context)
	ListFollowing(c *gin.Context)
	RegisterRoutes(r *gin.RouterGroup)
}

func (controller userController) ListUsers(c *gin.Context) {
	searchParam := c.Query("searchParam")
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		offset = 0
	}

	transactionId := c.GetHeader("transactionId")

	result, errorObj := controller.service.ListUsers(searchParam, limit, offset, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.JSON(200, result)
	}
}

func (controller userController) GetUser(c *gin.Context) {
	userId := c.Param("userId")
	transactionId := c.GetHeader("transactionId")

	result, errorObj := controller.service.GetUser(userId, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.JSON(200, result)
	}
}

func (controller userController) CreateUser(c *gin.Context) {
	var user models.User
	transactionId := c.GetHeader("transactionId")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, httpModels.ErrorResponse{ErrorCode: httpModels.ECODE_VALIDATION_ERROR, ErrorMessage: err.Error(), TransactionId: transactionId})
		return
	}

	errorObj := controller.service.CreateUser(&user, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.JSON(http.StatusCreated, user)
	}
}

func (controller userController) UpdateUser(c *gin.Context) {
	var user models.User
	userId := c.Param("userId")
	transactionId := c.GetHeader("transactionId")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, httpModels.ErrorResponse{ErrorCode: httpModels.ECODE_VALIDATION_ERROR, ErrorMessage: err.Error(), TransactionId: transactionId})
		return
	}

	if user.ID != "" && user.ID != userId {
		c.JSON(400, httpModels.ErrorResponse{ErrorCode: httpModels.ECODE_VALIDATION_ERROR, ErrorMessage: "userId in path and body should be same", TransactionId: transactionId})
		return
	}

	errorObj := controller.service.UpdateUser(&user, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func (controller userController) DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	transactionId := c.GetHeader("transactionId")

	errorObj := controller.service.DeleteUser(userId, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func (controller userController) FollowUser(c *gin.Context) {
	userId := c.Param("userId")
	followedUserId := c.Param("followedUserId")
	transactionId := c.GetHeader("transactionId")

	errorObj := controller.service.FollowUser(userId, followedUserId, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func (controller userController) UnFollowUser(c *gin.Context) {
	userId := c.Param("userId")
	followedUserId := c.Param("followedUserId")
	transactionId := c.GetHeader("transactionId")

	errorObj := controller.service.UnFollowUser(userId, followedUserId, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

func (controller userController) ListFollowers(c *gin.Context) {
	userId := c.Param("userId")
	searchParam := c.Query("searchParam")
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		offset = 0
	}
	transactionId := c.GetHeader("transactionId")

	result, errorObj := controller.service.ListFollowers(userId, searchParam, limit, offset, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.JSON(200, result)
	}
}

func (controller userController) ListFollowing(c *gin.Context) {
	userId := c.Param("userId")
	searchParam := c.Query("searchParam")
	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		offset = 0
	}
	transactionId := c.GetHeader("transactionId")

	result, errorObj := controller.service.ListFollowing(userId, searchParam, limit, offset, transactionId)

	if errorObj != nil {
		c.JSON(errorObj.StatusCode, errorObj)
		return
	} else {
		c.JSON(200, result)
	}
}

func (controller userController) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("", controller.ListUsers)
	router.GET("/:userId", controller.GetUser)
	router.POST("", controller.CreateUser)
	router.PUT("/:userId", controller.UpdateUser)
	router.DELETE("/:userId", controller.DeleteUser)
	router.GET("/:userId/followers", controller.ListFollowers)
	router.GET("/:userId/following", controller.ListFollowing)
	router.POST("/:userId/follow/:followedUserId", controller.FollowUser)
	router.DELETE("/:userId/follow/:followedUserId", controller.UnFollowUser)
}

func NewUserController(service service.UserService, logger *zap.SugaredLogger) UserController {
	return &userController{service, logger}
}
