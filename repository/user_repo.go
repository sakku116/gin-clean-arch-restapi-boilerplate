package repository

import (
	"github.com/qiniu/qmgo"
)

type UserRepo struct {
	coll *qmgo.Collection
}

type IUserRepo interface {
}

func NewUserRepo(coll *qmgo.Collection) IUserRepo {
	return &UserRepo{
		coll: coll,
	}
}
