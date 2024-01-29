package param

import "fiber-crud/entity"

type UserParam struct {
	PageNo   int
	PageSize int
	User     *entity.User
}
