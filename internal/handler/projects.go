package handler

import (
	"GoProject/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createProjects(c *gin.Context) {
	var input models.Projects
	if err := c.Bind(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	err := h.services.Projects.CreateProject(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Project created successfully"})
}

func (h *Handler) idProjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("handler - idProjects - invalid id param - %s", err.Error()))
		return
	}
	project, err := h.services.Projects.GetProjectByID(id)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, project)

}

func (h *Handler) idUserProjects(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("handler - idUserProjects - invalid id param - %s", err.Error()))
		return
	}
	users, err := h.services.Projects.GetUsersByProjectID(id)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, users)

}

func (h *Handler) submitProjects(c *gin.Context) {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("handler - idUserProjects - invalid projectid param - %s", err.Error()))
		return
	}
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("handler - idUserProjects - invalid userid param - %s", err.Error()))
		return
	}
	err = h.services.Projects.SubmitApplication(projectID, userID)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Application submitted successfully",
	})
}
