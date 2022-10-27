package handler

import (
	"net/http"

	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"github.com/gin-gonic/gin"
)

var loginErr string = ""
var registerErr string = ""

func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	input.Email = c.Request.FormValue("email")
	input.Username = c.Request.FormValue("username")
	input.Password = c.Request.FormValue("password")

	_, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		registerErr = err.Error()
		c.Redirect(301, "/auth/register")
		return
	}

	c.Redirect(301, "/auth/login")
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) register(c *gin.Context) {
	var regBool bool
	if registerErr != "" {
		regBool = true
	} else {
		regBool = false
	}
	c.HTML(http.StatusOK, "register.page.tmpl", gin.H{
		"RegisterErr": regBool,
		"ErrMessage":  registerErr,
	})
}

func (h *Handler) login(c *gin.Context) {
	var logBool bool
	if loginErr != "" {
		logBool = true
	} else {
		logBool = false
	}
	c.HTML(http.StatusOK, "login.page.tmpl", gin.H{
		"LoginErr":   logBool,
		"ErrMessage": loginErr,
	})
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	input.Username = c.Request.FormValue("username")
	input.Password = c.Request.FormValue("password")

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		loginErr = err.Error()
		c.Redirect(301, "/auth/login")
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*24, "", "", false, true)
	c.Redirect(301, "/api/notes")
}

func (h *Handler) signOut(c *gin.Context) {
	c.SetCookie("Authorization", "", 0, "", "", false, true)
	c.Redirect(301, "/")
}
