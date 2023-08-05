package handlers

import (
	"freelance/clinic_queue/api/http"
	"freelance/clinic_queue/models"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateService(c *gin.Context) {
	var service models.CreateService

	err := c.ShouldBindJSON(&service)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Service().CreateService(
		c.Request.Context(),
		&service,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}


// Get List
func (h *Handler) GetListService(c *gin.Context) {
	offset, err := h.getOffsetParam(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	limit, err := h.getLimitParam(c)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Service().GetListService(
		c.Request.Context(),
		&models.GetListServiceReq{
			Limit:  int32(limit),
			Offset: int32(offset),
			ClinicId: c.Query("clinic_id"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Update
func (h *Handler) UpdateService(c *gin.Context) {
	var service models.UpdateService

   err := c.ShouldBindJSON(&service)
   if err != nil {
	   h.handleResponse(c, http.BadRequest, err.Error())
	   return
   }

   resp, err := h.db.Service().UpdateService(
	   c.Request.Context(),
	   &service,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.OK, resp)
}

// Delete
func (h *Handler) DeleteService(c *gin.Context) {
	Id := c.Param("id")
   resp, err := h.db.Service().DeleteService(
	   c.Request.Context(),
	   Id,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.OK, resp)
}