package controller

import (
	"github.com/gin-gonic/gin"
	hikarilibbackend "github.com/yuminekosan/hikariLibBackend"
	"github.com/yuminekosan/hikariLibBackend/internal"
	"net/http"
)

func (r *Routes) signUp(c *gin.Context) {
	var input hikarilibbackend.User

	if err := c.ShouldBindJSON(&input); err != nil {
		internal.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := r.services.Authorization.CreateUser(input)
	if err != nil {
		internal.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (r *Routes) signIn(c *gin.Context) {}
