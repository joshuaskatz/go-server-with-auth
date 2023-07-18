package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/register", middleware.EmailValidation, middleware.PasswordValidation, controllers.Register, middleware.VerificationEmail)
	r.Use(middleware.EmailValidation, middleware.IsVerified).POST("/login", controllers.Login)
}
