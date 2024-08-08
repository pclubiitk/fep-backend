package prof

import (
	"github.com/gin-gonic/gin"
	"github.com/pclubiitk/fep-backend/middleware"
)

func extractProfID(ctx *gin.Context) (uint, error) {
	user_email := middleware.GetUserID(ctx)
	return FetchProfIDByEmail(ctx, user_email)
}
