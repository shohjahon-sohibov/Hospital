package api

import (
	"freelance/clinic_queue/api/handlers"
	"freelance/clinic_queue/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter godoc
// @description This is a api gateway
func SetUpRouter(h handlers.Handler, cfg config.Config) (r *gin.Engine) {
	r = gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.Use(customCORSMiddleware())

	// BRANCH
	r.POST("/signup", h.SignUp)
	r.POST("/login", h.Login)

	// USER
	r.POST("/user", h.CreateUser)
	r.GET("/user/:id", h.GetSingleUser)
	r.GET("/user", h.GetListUser)
	r.PUT("/user", h.UpdateUser)
	r.DELETE("/user/:id", h.DeleteUser)

	// DIAGNOSIS
	r.POST("/diagnosis", h.CreateDiagnosis)
	r.GET("/diagnosis/:id", h.GetSingleDiagnosis)
	r.GET("/diagnosis-list", h.GetListDiagnosis)
	r.PUT("/diagnosis", h.UpdateDiagnosis)
	r.DELETE("/diagnosis/:id", h.DeleteDiagnosis)

	// HOSPITAL
	r.POST("/hospital", h.CreateHospital)
	r.GET("/hospital/:id", h.GetSingleHospital)
	r.GET("/hospital", h.GetListHospital)
	r.PUT("/hospital", h.UpdateHospital)
	r.DELETE("/hospital/:id", h.DeleteHospital)

	// SERVICE
	r.POST("/service", h.CreateService)
	r.GET("/service", h.GetListService)
	r.PUT("/service", h.UpdateService)
	r.DELETE("/service/:id", h.DeleteService)

	// QUEUE
	r.POST("/queue", h.CreateQueue)
	r.GET("/queue", h.GetListQueue)
	r.PUT("/queue", h.ChangeStatusQueue)
	r.DELETE("/queue/:id", h.DeleteQueue)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
