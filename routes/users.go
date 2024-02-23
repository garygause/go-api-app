package routes

import (
	"net/http"
	"strconv"

	"github.com/garygause/go-api-app/models"
	"github.com/gin-gonic/gin"
)

func getUserById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User id is invalid."})
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve User."})
		return
	}
	context.JSON(http.StatusOK, user)
}

func getUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get Users."})
		return
	}
	context.JSON(http.StatusOK, users)
}

func createUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save User."})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "User created.", "user": user})
}

func updateUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User id is invalid."})
		return
	}
	_, err = models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve User."})
		return
	}

	var u models.User
	err = context.ShouldBindJSON(&u)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	u.ID = id
	err = u.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update User."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User updated."})
}


func deleteUser(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "User id is invalid."})
		return
	}
	user, err := models.GetUserById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve User."})
		return
	}
	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete User."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User deleted."})
}