package repository

import (
	"context"
	"database/sql"
	"fiber-crud/entity"
	"fiber-crud/param"
)

type UserRepository interface {
	Create(ctx context.Context, db *sql.Tx, user entity.User) (entity.User, error)
	Get(ctx context.Context, db *sql.Tx, param *param.UserParam) ([]entity.User, error)
	GetById(ctx context.Context, db *sql.DB, user entity.User) (entity.User, error)
	Update(ctx context.Context, db *sql.DB, user entity.User) (entity.User, error)
	Delete(ctx context.Context, db *sql.DB, user entity.User) (entity.User, error)
}
