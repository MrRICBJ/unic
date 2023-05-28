package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	userRole            = "userRole"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "token is empty")
		return
	}

	userId, role, err := h.service.ParseToken(headerParts[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	c.Set(userRole, role)
}

//func getUserId(c *gin.Context) (int, error) {
//	id, ok := c.Get(userCtx)
//	if !ok {
//		return 0, errors.New("user id not found")
//	}
//
//	idInt, ok := id.(int)
//	if !ok {
//		return 0, errors.New("user id is of invalid type")
//	}
//
//	return idInt, nil
//}
