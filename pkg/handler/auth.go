package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	todo "todo-app"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateUser(&input)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) signIn(c *gin.Context) {

}
