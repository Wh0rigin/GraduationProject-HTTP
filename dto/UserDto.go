package dto

import "github.com/Wh0rigin/GraduationProject/po"

type UserDto struct {
	Name      string `Json:"name"`
	Telephone string `Json:"telephone"`
	Manager   bool   `Json:"manager"`
}

func NewUserDto(user po.User) *UserDto {
	return &UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
		Manager:   user.Manager,
	}
}

func NewUserDtoByReference(user *po.User) *UserDto {
	return &UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
		Manager:   user.Manager,
	}
}
