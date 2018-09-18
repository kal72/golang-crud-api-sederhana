package userservice

import (
	"golang-crud-api-sederhana/repository"
	"golang-crud-api-sederhana/model"
)

/** business logic */

var userRepo repository.UserRepository = &repository.UserRepositoryImpl{}

func Find() []*model.User {
	return userRepo.Find()
}

func FindById(id int) *model.User{
	return userRepo.FindById(id)
}

func Save(user *model.User) *model.User {
	return userRepo.Save(user)
}

func Delete(id int) string{
	return userRepo.Delete(id)
}
