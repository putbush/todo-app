package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app"
)

func (h *Handler) getAllItems(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	items, err := h.services.GetAllItems(userID, listID)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.GetItemByID(userID, itemID)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) createItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.Item
	if err = c.BindJSON(&input); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoItem.CreateItem(input, listID, userID)
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.UpdateItemInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	item, err := h.services.TodoItem.Update(userID, itemID, input)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) deleteItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	itemID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.DeleteItemByID(itemID, userID); err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
