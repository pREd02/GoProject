package handler

import (
	"GoProject/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrganDataId struct {
	OrganId int `json:"organ_id"`
}

func (h *Handler) registerOrganization(c *gin.Context) {

	var input models.Organization
	if err := c.Bind(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	organizationID, err := h.services.Organization.Create(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, OrganDataId{
		OrganId: organizationID,
	})
}

func (h *Handler) idOrganization(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, fmt.Errorf("handler - idOrganization - invalid id param - %s", err.Error()))
		return
	}
	organization, err := h.services.Organization.GetOrganById(id)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, organization)
}
