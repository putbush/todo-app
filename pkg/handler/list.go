package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	todo "todo-app"
)

type getAllListsResponse struct {
	Data []todo.TodoList `json:"data"`
}

// @Summary Get All Lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllListsResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllLists(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	lists, err := h.services.GetAllLists(userID)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

func (h *Handler) getListByID(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.GetListByID(listID, userID)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags lists
// @Description create todo list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body todo.TodoList true "list info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateList(input, userID)
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) updateList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	var input todo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = input.Validate(); err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}
	newList, err := h.services.TodoList.Update(listID, userID, input)
	if err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, newList)
}

func (h *Handler) deleteList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResonse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err = h.services.DeleteListByID(listID, userID); err != nil {
		newErrorResonse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{Status: "ok"})
}
