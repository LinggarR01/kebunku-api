package handlers

import (
	"fmt"
	"kebunku-api/dto"
	"net/http"
	"path/filepath"
	"time"

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

func (h *compHandlers) GetTanaman(c *gin.Context) {
	data, err := h.service.GetTanaman()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.Response{Body: data})
}

func (h *compHandlers) UploadTanaman(c *gin.Context) {
	var file_url *string

	tanaman_id := c.Query("id")
	if tanaman_id == "" {
            c.JSON(http.StatusBadRequest, dto.Response{ Error: "Tanaman ID is required"})
            return
    }

	err := c.Request.ParseMultipartForm(10 << 20)
	if err != nil {
			c.JSON(http.StatusBadRequest, dto.Response{ Error: err.Error()})
			return
	}

	_, header, err := c.Request.FormFile("file")
	if err != nil {
			c.JSON(http.StatusBadRequest, dto.Response{ Error: "Bad Request"})
			return
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(header.Filename))

	filePath := filepath.Join("public", "tanaman", filename)
	err = c.SaveUploadedFile(header, filePath)
	if err != nil {
			c.JSON(http.StatusInternalServerError, dto.Response{ Error: "Failed to save file"})
			return
	} else {
			file_url_new := "/tanaman/" + filename

			if file_url_new != "" {
					file_url = &file_url_new
			}
	}

	err = h.service.UploadTanaman(*file_url, tanaman_id)
	if err!= nil {
            c.JSON(http.StatusInternalServerError, dto.Response{ Error: err.Error()})
            return
    }

	c.JSON(http.StatusOK, dto.Response{Message: "Product Request file successfully saved!", Body: file_url})
}