package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"rest-api/pkg/models"
	"rest-api/pkg/repository/postgres"
)

type Handler struct {
	userRepo *postgres.UserRepository
}

type GetUserUri struct {
	UserId string `uri:"userId" binding:"required,uuid"`
}

type CreateUser struct {
	Login string `json:"login" binding:"required"`
}

type UpdateUser struct {
	Login string `json:"login" binding:"required"`
}

func NewHandler(ur *postgres.UserRepository) *Handler {
	return &Handler{
		userRepo: ur,
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	//validate uri
	uri := GetUserUri{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	userId := uri.UserId

	//find user in db
	user, err := h.userRepo.GetUserByUuid(c, userId)
	if err != nil {
		if fmt.Sprintf("%v", err) == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"status": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err})
		}
		return
	}

	log.Printf("Get_User_Result, err: %v, %v", user, err)

	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"uuid":       user.Uuid,
		"login":      user.Login,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
		"deleted_at": user.DeletedAt,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {
	//validate body
	var json CreateUser
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	userLogin := json.Login

	//check if user exists in db
	_, err := h.userRepo.GetUserByLogin(c, userLogin)
	if err != nil {
		if fmt.Sprintf("%v", err) == "record not found" {
			//creating new user
			id := uuid.New()
			newUser := models.User{}
			newUser.Uuid = fmt.Sprintf("%v", id)
			newUser.Login = userLogin
			user, err := h.userRepo.CreateUser(c, &newUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
				return
			}

			log.Printf("Create_User_Result, err: %v, %v", user, err)
			c.JSON(http.StatusCreated, gin.H{
				"id":         user.ID,
				"uuid":       user.Uuid,
				"login":      user.Login,
				"created_at": user.CreatedAt,
				"updated_at": user.UpdatedAt,
				"deleted_at": user.DeletedAt,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err})
		}
		return
	}

	c.JSON(http.StatusForbidden, gin.H{"status": fmt.Sprintf("user with login '%v' already exists", userLogin)})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	//validate uri
	uri := GetUserUri{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	userId := uri.UserId

	//validate body
	var json UpdateUser
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	userLogin := json.Login

	//check if user exists in db
	_, err := h.userRepo.GetUserByUuid(c, userId)
	if err != nil {
		if fmt.Sprintf("%v", err) == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"status": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err})
		}
		return
	}

	//update user
	updateUser := models.User{}
	updateUser.Uuid = fmt.Sprintf("%v", userId)
	updateUser.Login = userLogin
	updateUserResult, err := h.userRepo.UpdateUsersLogin(c, &updateUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	log.Printf("Update_User_Result, err: %v, %v", updateUserResult, err)
	c.JSON(http.StatusOK, gin.H{
		"uuid":  updateUserResult.Uuid,
		"login": updateUserResult.Login,
	})

}

func (h *Handler) DeleteUser(c *gin.Context) {
	//validate uri
	uri := GetUserUri{}
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": err.Error()})
		return
	}
	userId := uri.UserId

	//check if user exists in db
	_, err := h.userRepo.GetUserByUuid(c, userId)
	if err != nil {
		if fmt.Sprintf("%v", err) == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"status": "user not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"status": err})
		}
		return
	}

	//delete user
	deleteUser := models.User{}
	deleteUser.Uuid = fmt.Sprintf("%v", userId)
	err = h.userRepo.DeleteUser(c, &deleteUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
