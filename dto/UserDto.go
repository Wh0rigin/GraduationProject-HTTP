package dto

import "github.com/Wh0rigin/GraduationProject/bean"

type UserDto struct {
	Name      string `Json:"name"`
	Telephone string `Json:"telephone"`
}

func NewUserDto(user bean.User) *UserDto {
	return &UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

func NewUserDtoByReference(user *bean.User) *UserDto {
	return &UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
