package bean

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string //`gorm:"type:varchar(20);not null"`
	Telephone string // `gorm:"type:varchar(11);not null"`
	Password  string //`gorm:"type:size:255;not null"`
	Version   int    //`gorm:"type:Integer;not null"`
	IsDestory bool   // `gorm:"type:BOOLEAN;not null"`
}

//constructor
func NewUser(n string, t string, p string) *User {
	return &User{
		Name:      n,
		Telephone: t,
		Password:  p,
		Version:   1,
		IsDestory: false,
	}
}

//get set
// func (self *User) GetName() string {
// 	return self.name
// }

// func (self *User) GetTelephone() string {
// 	return self.telephone
// }

// func (self *User) GetPassword() string {
// 	return self.password
// }

// func (self *User) GetIsDestory() bool {
// 	return self.isDestory
// }

// func (self *User) SetName(value string) {
// 	self.name = value
// }

// func (self *User) SetTelephone(value string) {
// 	self.telephone = value
// }

// func (self *User) SetPassword(value string) {
// 	self.password = value
// }
// func (self *User) SetIsDestory() {
// 	self.isDestory = true
// }
