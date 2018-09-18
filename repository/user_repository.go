package repository

import "golang-crud-api-sederhana/model"

/**
* interface repository for database operation
 */
type UserRepository interface {

	Find() []*model.User

	FindById(id int) *model.User

	Save(user *model.User) *model.User

	Delete(id int) string
}
