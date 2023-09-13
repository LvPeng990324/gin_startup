package dto

import "contract_system_server/model"

type UserDto struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Id:    user.Id,
		Name:  user.Name,
		Phone: user.Phone,
	}
}
