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

func NewUserService(
	db *sql.DB,
	userRepository repository.UserRepository,
) *UserServiceImpl {
	return &UserServiceImpl{
		DB:             db,
		UserRepository: userRepository,
	}
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
	tx, err := service.DB.Begin()
	if err != nil {
		return []result.UserResult{}, err
	}
	users, err := service.UserRepository.Get(
		ctx,
		tx,
		&param,
	)
	return helper.UserEntitiesToUserResults(users), err
}

func (service *UserServiceImpl) GetById(ctx context.Context, user entity.User) (result.UserResult, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return result.UserResult{}, err
	}
	selected, err := service.UserRepository.GetById(
		ctx,
		tx,
		user,
	)
	return helper.UserEntityToUserResult(selected), err
}

func (service *UserServiceImpl) Update(ctx context.Context, user entity.User) (result.UserResult, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return result.UserResult{}, err
	}
	current, err := service.UserRepository.GetById(
		ctx,
		tx,
		user,
	)
	if user.Name != "" {
		current.Name = user.Name
	}
	if user.PhoneNumber != "" {
		current.PhoneNumber = user.PhoneNumber
	}
	updated, err := service.UserRepository.Update(
		ctx,
		tx,
		current,
	)
	return helper.UserEntityToUserResult(updated), err
}

func (service *UserServiceImpl) Delete(ctx context.Context, user entity.User) (result.UserResult, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return result.UserResult{}, err
	}
	deleted, err := service.UserRepository.Delete(
		ctx,
		tx,
		user,
	)
	return helper.UserEntityToUserResult(deleted), err
}
