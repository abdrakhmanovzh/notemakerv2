package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) userIdentity(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.Redirect(301, "/auth/login")
		return
	}

	if tokenString == "" {
		c.Redirect(301, "/auth/login")
		return
	}

	uid, username, err := h.services.Authorization.ParseToken(tokenString)
	if err != nil {
		c.Redirect(301, "/auth/login")
		return
	}

	c.Set("uid", uid)
	c.Set("username", username)
}
