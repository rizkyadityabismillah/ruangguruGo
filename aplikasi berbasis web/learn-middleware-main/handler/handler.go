package handler

import (
	"example-middleware/model"
	"example-middleware/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IService interface {
	Get() []model.Data
	GetById(id int) model.Data
	Update(model.Data) error
	Delete(id int) error
	Create(model.Data) error
}

type IHandler interface {
	CreateData(c *gin.Context)
	GetData(c *gin.Context)
	// GetDataById(c *gin.Context)
	// UpdateData(c *gin.Context)
	// DeleteData(c *gin.Context)
}

type Handler struct {
	Service service.IService
}

func NewHandler(service service.IService) IHandler {
	return &Handler{
		Service: service,
	}
}

// pengecekan
// username sama password
// admin, 1234

func (h *Handler) CreateData(c *gin.Context) {
	// ngedapetin data
	var data model.Data

	// request
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.Service.Create(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "success create data",
	})
}

func (h *Handler) GetData(c *gin.Context) {
	// ngedapetin data
	datas := h.Service.Get()

	c.JSON(http.StatusOK, datas)
}
