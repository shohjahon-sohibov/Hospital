package handlers

import (
	"fmt"
	"freelance/clinic_queue/api/http"
	"freelance/clinic_queue/models"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateQueue(c *gin.Context) {
	var queue models.CreateQueue
	err := c.ShouldBindJSON(&queue)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Queue().CreateQueue(
		c.Request.Context(),
		&queue,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// Update
func (h *Handler) ChangeStatusQueue(c *gin.Context) {
	var status models.ChangeStatusQueue

	err := c.ShouldBindJSON(&status)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.Queue().ChangeStatusQueue(
		c.Request.Context(),
		&status,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Get Single
// func (h *Handler) GetSingleHospital(c *gin.Context) {
// 	Id := c.Param("id")
// 	resp, err := h.db.Hospital().GetSingleHospital(
// 		c.Request.Context(),
// 		Id,
// 	)
// 	if err != nil {
// 		h.handleResponse(c, http.GRPCError, err.Error())
// 		return
// 	}

// 	h.handleResponse(c, http.OK, resp)
// }

// Get List
func (h *Handler) GetListQueue(c *gin.Context) {
	fmt.Println(1)
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

	resp, err := h.db.Queue().GetListQueue(
		c.Request.Context(),
		&models.GetListQueueReq{
			Limit:  int32(limit),
			Offset: int32(offset),
			DoctorId: c.Query("doctor_id"),
			Date: c.Query("date"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Delete
func (h *Handler) DeleteQueue(c *gin.Context) {
	Id := c.Param("id")
   resp, err := h.db.Queue().DeleteQueue(
	   c.Request.Context(),
	   Id,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.OK, resp)
}