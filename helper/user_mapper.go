package helper

import (
	"fiber-crud/entity"
	"fiber-crud/result"
)

func UserEntityToUserResult(entity entity.User) result.UserResult {
	return result.UserResult{
		ID:          uint(entity.ID),
		Name:        entity.Name,
		PhoneNumber: entity.PhoneNumber,
	}
}

func UserEntitiesToUserResults(entities []entity.User) []result.UserResult {
	results := []result.UserResult{}
	for _, entity := range entities {
		results = append(results, UserEntityToUserResult(entity))
	}
	return results
}
