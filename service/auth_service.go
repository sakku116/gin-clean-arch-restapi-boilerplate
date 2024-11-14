package service

import (
	"restapi/repository"
)

type AuthService struct {
	userRepo repository.IUserRepo
}

type IAuthService interface {
}

func NewAuthService(userRepo repository.IUserRepo) IAuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}
