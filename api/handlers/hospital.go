package handlers

import (
	"freelance/clinic_queue/api/http"
	"freelance/clinic_queue/models"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateHospital(c *gin.Context) {
	var hospital models.CreateHospital

	err := c.ShouldBindJSON(&hospital)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Hospital().CreateHospital(
		c.Request.Context(),
		&hospital,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// Update
func (h *Handler) UpdateHospital(c *gin.Context) {
	var hospital models.UpdateHospital

	err := c.ShouldBindJSON(&hospital)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Hospital().UpdateHospital(
		c.Request.Context(),
		&hospital,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get Single
func (h *Handler) GetSingleHospital(c *gin.Context) {
	Id := c.Param("id")
	resp, err := h.db.Hospital().GetSingleHospital(
		c.Request.Context(),
		Id,
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get List
func (h *Handler) GetListHospital(c *gin.Context) {
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

	resp, err := h.db.Hospital().GetListHospital(
		c.Request.Context(),
		&models.GetListHospitalReq{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Delete
func (h *Handler) DeleteHospital(c *gin.Context) {
	Id := c.Param("id")
   resp, err := h.db.Hospital().DeleteHospital(
	   c.Request.Context(),
	   Id,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.OK, resp)
}