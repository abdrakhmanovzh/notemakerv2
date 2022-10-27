package handler

import (
	"net/http"
	"strconv"

	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createNote(c *gin.Context) {
	c.HTML(http.StatusOK, "create.page.tmpl", nil)
}

func (h *Handler) createdNote(c *gin.Context) {
	userId, ok := c.Get("uid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}

	var input model.Note

	input.Title = c.Request.FormValue("title")
	input.Content = c.Request.FormValue("content")

	_, err := h.services.Note.Create(userId.(int), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redirect(301, "/api/notes")
}

func (h *Handler) getAllNotes(c *gin.Context) {
	var toCheck bool
	userId, ok := c.Get("uid")
	if !ok {
		c.Redirect(301, "/auth/login")
		return
	}

	username, ok := c.Get("username")
	if !ok {
		c.Redirect(301, "/auth/login")
		return
	}

	notes, err := h.services.Note.GetAll(userId.(int))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(notes) == 0 {
		toCheck = false
	} else {
		toCheck = true
	}

	c.HTML(http.StatusOK, "notes.page.tmpl", gin.H{
		"Username": username,
		"ToCheck":  toCheck,
		"Note":     notes,
	})
}

func (h *Handler) getNoteById(c *gin.Context) {
	userId, ok := c.Get("uid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	note, err := h.services.Note.GetById(userId.(int), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "show.page.tmpl", gin.H{
		"Title":   note.Title,
		"Content": note.Content,
		"ID":      id,
	})
}

func (h *Handler) updatedNote(c *gin.Context) {
	userId, ok := c.Get("uid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	var input model.UpdateNoteInput

	input.Title = c.Request.FormValue("title")
	input.Content = c.Request.FormValue("content")

	h.services.Update(userId.(int), id, input)
	c.Redirect(301, "/api/notes")
}

func (h *Handler) updateNote(c *gin.Context) {
	userId, ok := c.Get("uid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	note, err := h.services.Note.GetById(userId.(int), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "update.page.tmpl", gin.H{
		"Title":   note.Title,
		"Content": note.Content,
		"ID":      id,
	})
}

func (h *Handler) deleteNote(c *gin.Context) {
	userId, ok := c.Get("uid")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user id not found"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	del := h.services.Note.Delete(userId.(int), id)
	if del != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": del.Error()})
		return
	}

	c.Redirect(301, "/api/notes")
}
