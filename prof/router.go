package prof

import (
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.Engine) {
	admin := r.Group("/api/admin/prof")
	{
		admin.GET("/hello", greetingHandler)
		admin.GET("", getAllProfHandler)
		admin.GET("/:pid", getProfHandler)
		admin.GET("/limited", getLimitedProfHandler)

		admin.PUT("", updateProfHandler)
		admin.POST("", addNewHandler)

		admin.DELETE("/:cid", deleteProfHandler)

	}
}

func ProfRouter(r *gin.Engine) {
	prof := r.Group("/api/prof")
	{
		prof.GET("", getProfHandler)
		prof.GET("/hello", greetingHandler) // to be updated

	}
}
