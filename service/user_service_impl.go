package service

import (
	"context"
	"database/sql"
	"fiber-crud/entity"
	"fiber-crud/helper"
	"fiber-crud/param"
	"fiber-crud/repository"
	"fiber-crud/result"
)

type UserServiceImpl struct {
	DB             *sql.DB
	UserRepository repository.UserRepository
}

func (service *UserServiceImpl) Create(ctx context.Context, user entity.User) (result.UserResult, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return result.UserResult{}, err
	}
	inserted, err := service.UserRepository.Create(
		ctx,
		tx,
		user,
	)
	return helper.UserEntityToUserResult(inserted), err
}

func (service *UserServiceImpl) Get(ctx context.Context, param param.UserParam) ([]result.UserResult, error) {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) GetById(ctx context.Context, user entity.User) (result.UserResult, error) {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) Update(ctx context.Context, user entity.User) (result.UserResult, error) {
	panic("not implemented") // TODO: Implement
}

func (service *UserServiceImpl) Delete(ctx context.Context, user entity.User) (result.UserResult, error) {
	panic("not implemented") // TODO: Implement
}
