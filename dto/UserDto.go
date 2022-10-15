package dto

import "github.com/Wh0rigin/GraduationProject/po"

type UserDto struct {
	Name      string `Json:"name"`
	Telephone string `Json:"telephone"`
}

func NewUserDto(user po.User) *UserDto {
	return &UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

func NewUserDtoByReference(user *po.User) *UserDto {
	return &UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
