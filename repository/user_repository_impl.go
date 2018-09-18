package repository

import (
	"golang-crud-api-sederhana/model"
	"fmt"
)

/**
* implementation from interface repository for database operation
 */

type UserRepositoryImpl struct {
}


func (r *UserRepositoryImpl) Find() []*model.User {
	var users []*model.User

	model.GetOrmObject().QueryTable("user").All(&users)

	return users
}

func (r *UserRepositoryImpl) FindById(id int) *model.User {
	user := model.User{Id: id}
	err := model.GetOrmObject().Read(&user)
	if err != nil {
		return nil
	}

	return &user
}

func (r *UserRepositoryImpl) Save(user *model.User) *model.User {
	model.GetOrmObject().Insert(user)

	return user
}

func (r *UserRepositoryImpl) Delete(id int) string {
	user := model.User{Id: id}
	if num, err := model.GetOrmObject().Delete(&user); err != nil {
		fmt.Println(num)
		return "failed"
	}

	fmt.Println("deleted")
	return "deleted"
}
