package service

import (
	"context"
	"fiber-crud/entity"
	"fiber-crud/param"
	"fiber-crud/result"
)

type UserService interface {
	Create(ctx context.Context, user entity.User) (result.UserResult, error)
	Get(ctx context.Context, param param.UserParam) ([]result.UserResult, error)
	GetById(ctx context.Context, user entity.User) (result.UserResult, error)
	Update(ctx context.Context, user entity.User) (result.UserResult, error)
	Delete(ctx context.Context, user entity.User) (result.UserResult, error)
}
