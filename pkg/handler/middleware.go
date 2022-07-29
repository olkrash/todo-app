package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userID"
)

//func (h *Handler) UserIdentity(c *gin.Context) {
//	header := c.GetHeader(authorizationHeader)
//	if header == "" {
//		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
//		return
//	}
//
//	headerParts := strings.Split(header, "")
//	if len(headerParts) != 2 {
//		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
//		return
//	}
//
//	userId, err := h.services.Authorization.ParseToken(headerParts[1])
//	if err != nil {
//		newErrorResponse(c, http.StatusUnauthorized, err.Error())
//		return
//	}
//	c.Set(userCtx, userId)
//
//}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id ot found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is not valid of type")
		return 0, errors.New("user id not found")
	}
	return idInt, nil
}
