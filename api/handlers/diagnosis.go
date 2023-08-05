package handlers

import (
	"freelance/clinic_queue/api/http"
	"freelance/clinic_queue/models"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateDiagnosis(c *gin.Context) {
	var diagnosis models.CreateDiagnosis

   err := c.ShouldBindJSON(&diagnosis)
   if err != nil {
	   h.handleResponse(c, http.BadRequest, err.Error())
	   return
   }

   resp, err := h.db.Diagnosis().CreateDiagnosis(
	   c.Request.Context(),
	   &diagnosis,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.Created, resp)
}

// Get Single
func (h *Handler) GetSingleDiagnosis(c *gin.Context) {
	Id := c.Param("id")
   resp, err := h.db.Diagnosis().GetSingleDiagnosis(
	   c.Request.Context(),
	   Id,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.Created, resp)
}

// Get List
func (h *Handler) GetListDiagnosis(c *gin.Context) {
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

	resp, err := h.db.Diagnosis().GetListDiagnosis(
		c.Request.Context(),
		&models.GetListDiagnosisReq{
			Limit:  int32(limit),
			Offset: int32(offset),
			Date: c.Query("date"),
			UserId: c.Query("user_id"),
			DoctorId: c.Query("doctor_id"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// Update
func (h *Handler) UpdateDiagnosis(c *gin.Context) {
	var diagnosis models.UpdateDiagnosis

   err := c.ShouldBindJSON(&diagnosis)
   if err != nil {
	   h.handleResponse(c, http.BadRequest, err.Error())
	   return
   }

   resp, err := h.db.Diagnosis().UpdateDiagnosis(
	   c.Request.Context(),
	   &diagnosis,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.Created, resp)
}

// Delete
func (h *Handler) DeleteDiagnosis(c *gin.Context) {
	Id := c.Param("id")
   resp, err := h.db.Diagnosis().DeleteDiagnosis(
	   c.Request.Context(),
	   Id,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.Created, resp)
}