package handlers

import (
	"freelance/clinic_queue/api/http"
	"freelance/clinic_queue/config"
	"freelance/clinic_queue/storage"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/saidamir98/udevs_pkg/logger"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg config.Config
	log logger.LoggerI
	db  storage.StorageI
}

func NewHandler(cfg config.Config, log logger.LoggerI, storage storage.StorageI) Handler {
	return Handler{
		cfg: cfg,
		log: log,
		db:  storage,
	}
}

func (h *Handler) handleResponse(c *gin.Context, status http.Status, data interface{}) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"---Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	case code < 400:
		h.log.Warn(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	default:
		h.log.Error(
			"!!!Response--->",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
			logger.Any("data", data),
		)
	}

	c.JSON(status.Code, http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *Handler) getOffsetParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("offset", h.cfg.DefaultOffset)
	return strconv.Atoi(offsetStr)
}

func (h *Handler) getLimitParam(c *gin.Context) (offset int, err error) {
	offsetStr := c.DefaultQuery("limit", h.cfg.DefaultLimit)
	return strconv.Atoi(offsetStr)
}

func ASCIToChar(number int32) string {
	return string(number)
}

func WriteToExcel(f *excelize.File, data []interface{}, row int) {
	cellNum := 65
	for i := 0; i < len(data); i++ {
		cell := ASCIToChar(int32(cellNum)) + strconv.Itoa(row)
		f.SetCellValue("Sheet1", cell, data[i])

		cellNum = cellNum + 1
	}

}
