package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/mail"
)

func Router(mail_channel chan mail.Mail, r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.GET("/hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "hello"})
		})
		auth.POST("/login", loginHandler)
		auth.POST("/otp", otpHandler(mail_channel))
		auth.POST("/signup", signUpHandler(mail_channel))
		auth.GET("/whoami", whoamiHandler) // who am i, if not exploited

		auth.GET("/admins", getAllAdminDetailsHandler)
		auth.GET("/admins/:userID", getAdminDetailsHandler)
		auth.PUT("/admins/:userID/role", updateUserRole)
		auth.PUT("/admins/:userID/active", updateUserActiveStatus)
		auth.POST("/reset-password", resetPasswordHandler(mail_channel))

		auth.POST("/prof-signup", profSignUpHandler(mail_channel))

		auth.POST("/god/signup", godSignUpHandler(mail_channel))
		auth.POST("/god/login", godLoginHandler)
		auth.POST("/god/reset-password", godResetPasswordHandler(mail_channel))
	}
}
