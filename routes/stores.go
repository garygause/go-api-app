package routes

import (
	"net/http"
	"strconv"

	"github.com/garygause/go-api-app/models"
	"github.com/gin-gonic/gin"
)

func getStoreById(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Store id is invalid."})
		return
	}
	store, err := models.GetStoreById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve Store."})
		return
	}
	context.JSON(http.StatusOK, store)
}

func getStores(context *gin.Context) {
	stores, err := models.GetAllStores()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get Stores."})
		return
	}
	context.JSON(http.StatusOK, stores)
}

func createStore(context *gin.Context) {
	userId := context.GetInt64("userId")

	var store models.Store
	err := context.ShouldBindJSON(&store)
	
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	store.UserID = userId
	err = store.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save Store."})
		return
	}
	
	context.JSON(http.StatusCreated, gin.H{"message": "Store created.", "store": store})
}

func updateStore(context *gin.Context) {
	userId := context.GetInt64("userId")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Store id is invalid."})
		return
	}
	store, err := models.GetStoreById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve Store."})
		return
	}

	if store.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	var s models.Store
	err = context.ShouldBindJSON(&s)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	s.ID = id
	err = s.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update Store."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Store updated."})
}


func deleteStore(context *gin.Context) {
	userId := context.GetInt64("userId")

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Store id is invalid."})
		return
	}
	store, err := models.GetStoreById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not retrieve Store."})
		return
	}

	if store.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	err = store.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete Store."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Store deleted."})
}