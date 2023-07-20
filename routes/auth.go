package routes

import (
	"server/controllers"
	"server/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/register", middleware.EmailValidation, middleware.PasswordValidation, controllers.Register, middleware.VerificationEmail)
	r.POST("/login", middleware.EmailValidation, middleware.IsVerified, controllers.Login)
	r.POST("/verify/:jwt", controllers.Verify)
}
