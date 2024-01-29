package repository

import (
	"context"
	"database/sql"
	"errors"
	"fiber-crud/entity"
	"fiber-crud/param"
)

type UserRepositoryImpl struct {
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, db *sql.Tx, user entity.User) (entity.User, error) {
	result, err := db.ExecContext(
		ctx,
		"INSERT INTO user(name, phone_number) VALUES (?, ?)",
		user.Name,
		user.PhoneNumber,
	)
	if err != nil {
		return user, err
	}
	userID, err := result.LastInsertId()
	user.ID = int32(userID)
	return user, nil
}

func (repository *UserRepositoryImpl) Get(ctx context.Context, db *sql.Tx, param *param.UserParam) ([]entity.User, error) {
	rows, err := db.QueryContext(
		ctx,
		"SELECT id, name, phone_number from user LIMIT ? OFFSET ?",
		param.PageSize,
		(param.PageNo-1)*param.PageSize,
	)
	if err != nil {
		return []entity.User{}, err
	}
	results := []entity.User{}
	for rows.Next() {
		var id int32
		var name, phoneNumber string
		err := rows.Scan(&id, &name, &phoneNumber)
		results = append(results, entity.User{
			ID:          id,
			Name:        name,
			PhoneNumber: phoneNumber,
		})
		if err != nil {
			return results, err
		}
	}
	defer rows.Close()
	return results, nil
}

func (repository *UserRepositoryImpl) GetById(ctx context.Context, db *sql.DB, user entity.User) (entity.User, error) {
	row := db.QueryRowContext(
		ctx,
		"SELECT id, name, phone_number from user WHERE id = ? LIMIT 1",
		user.ID,
	)
	var id int32
	var name, phoneNumber string
	err := row.Scan(&id, &name, &phoneNumber)
	if err != nil {
		return entity.User{}, err
	}
	return entity.User{
		ID:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
	}, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, db *sql.DB, user entity.User) (entity.User, error) {
	result, err := db.ExecContext(
		ctx,
		"UPDATE user SET name = ?, phone_number = ? WHERE id = ?",
		user.Name,
		user.PhoneNumber,
		user.ID,
	)
	if err != nil {
		return user, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, err
	}
	if rows != 1 {
		return entity.User{}, errors.New("Unknown error. Update not successful")
	}
	return user, nil
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, db *sql.DB, user entity.User) (entity.User, error) {
	result, err := db.ExecContext(
		ctx,
		"DELETE from user WHERE id = ?",
		user.ID,
	)
	if err != nil {
		return user, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return entity.User{}, err
	}
	if rows != 1 {
		return entity.User{}, errors.New("Unknown error. Delete not successful")
	}
	return user, nil
}
