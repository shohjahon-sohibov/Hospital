package handlers

import (
	"fmt"
	"freelance/clinic_queue/api/http"
	"freelance/clinic_queue/models"

	"github.com/gin-gonic/gin"
)

// Create
func (h *Handler) CreateUser(c *gin.Context) {
 	var user models.CreateUser

	err := c.ShouldBindJSON(&user)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.db.User().CreateUser(
		c.Request.Context(),
		&user,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// Get Single
func (h *Handler) GetSingleUser(c *gin.Context) {
	Id := c.Param("id")
	resp, err := h.db.User().GetSingleUser(
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
func (h *Handler) GetListUser(c *gin.Context) {
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

	resp, err := h.db.User().GetListUser(
		c.Request.Context(),
		&models.GetUsersListReq{
			Limit:  int32(limit),
			Offset: int32(offset),
			Role: c.Query("role"),
			ClinicId: c.Query("clinic_id"),
			ServiceId: c.Query("service_id"),
			PhoneNumber: c.Query("phone_number"),
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.OK, resp)
}

// Update
func (h *Handler) UpdateUser(c *gin.Context) {
	var user models.UpdateUser

   err := c.ShouldBindJSON(&user)
   if err != nil {
	   h.handleResponse(c, http.BadRequest, err.Error())
	   return
   }

   resp, err := h.db.User().UpdateUser(
	   c.Request.Context(),
	   &user,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.OK, resp)
}

// Delete
func (h *Handler) DeleteUser(c *gin.Context) {
	Id := c.Param("id")
	fmt.Println(Id, 0)
   resp, err := h.db.User().DeleteUser(
	   c.Request.Context(),
	   Id,
   )

   if err != nil {
	   h.handleResponse(c, http.GRPCError, err.Error())
	   return
   }

   h.handleResponse(c, http.OK, resp)
}