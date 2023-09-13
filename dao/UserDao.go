package dao

import (
	"contract_system_server/common"
	"contract_system_server/model"
)

func CreateUser(user model.User) (err error) {
	DB := common.GetDB()
	sqlStr := "insert into user (name, phone, password) values (?, ?, ?)"
	_, err = DB.Exec(sqlStr, user.Name, user.Phone, user.Password)

	return err
}

func GetUserById(id uint) (user model.User, err error) {
	DB := common.GetDB()
	sqlStr := "select id, name, phone, password, is_deleted, created_at from user where id = ?"
	err = DB.Get(&user, sqlStr, id)

	return user, err
}

func GetAllUser() (users []model.User, err error) {
	DB := common.GetDB()
	sqlStr := "select id, name, phone, password, is_deleted, created_at from user"
	err = DB.Select(&users, sqlStr)

	return users, err
}
