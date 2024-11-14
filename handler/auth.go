package handler

import (
	"restapi/service"
	"restapi/utils/http_response"
)

type AuthHandler struct {
	respWriter  http_response.IResponseWriter
	authService service.IAuthService
}

type IAuthHandler interface {
}

func NewAuthHandler(respWriter http_response.IResponseWriter, authService service.IAuthService) AuthHandler {
	return AuthHandler{
		respWriter:  respWriter,
		authService: authService,
	}
}
