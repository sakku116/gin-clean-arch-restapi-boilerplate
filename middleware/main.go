package middleware

import (
	"restapi/service"
	"restapi/utils/http_response"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(respWriter http_response.IResponseWriter, authService service.IAuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		_ = c.GetHeader("Bearer")

		c.Next()
	}
}
