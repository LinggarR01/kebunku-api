package handlers

import (
	"kebunku-api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *compHandlers) RegisterTanaman(c *gin.Context) {
	// Implement the logic for registering a new tanaman
	var data dto.Tanaman
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{Error: "bad request"})
		return
	}

	err = h.service.RegisterTanaman(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Body: data})

}
