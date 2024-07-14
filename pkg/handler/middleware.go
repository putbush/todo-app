package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResonse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResonse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userID, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResonse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userID)
}
