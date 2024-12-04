package http_response

import (
	"restapi/domain/rest"
	error_utils "restapi/utils/error"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct{}

type IResponseWriter interface {
	HTTPCustomErr(ctx *gin.Context, err error)
	HTTPJsonErr(ctx *gin.Context, code int, message string, detail string, data interface{})
	HTTPJson(ctx *gin.Context, data interface{})
}

func NewResponseWriter() IResponseWriter {
	return &ResponseWriter{}
}

func (r *ResponseWriter) HTTPCustomErr(ctx *gin.Context, err error) {
	customErr, ok := err.(*error_utils.CustomErr)
	if ok {
		ctx.JSON(customErr.HttpCode, rest.BaseJSONResp{
			Code:    customErr.HttpCode,
			Message: customErr.Error(),
			Detail:  customErr.Detail,
			Data:    customErr.Data,
		})
		return
	}
	ctx.JSON(500, rest.BaseJSONResp{
		Code:    500,
		Message: "internal server error",
		Detail:  err.Error(),
		Data:    nil,
	})
}

func (r *ResponseWriter) HTTPJsonErr(ctx *gin.Context, code int, message string, detail string, data interface{}) {
	ctx.JSON(code, rest.BaseJSONResp{
		Code:    code,
		Message: message,
		Detail:  detail,
		Data:    data,
	})
}

func (r *ResponseWriter) HTTPJson(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, rest.BaseJSONResp{
		Code:    200,
		Message: "OK",
		Detail:  "",
		Data:    data,
	})
}
