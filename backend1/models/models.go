package models

type User struct{
	Id_user int `gorm:"primaryKey;autoIncrement:true"`
	Id_admin int 
	Nama string
	Password string
	Username string
	Is_active bool
}

type Admin struct{
	Id_admin int	`gorm:"primaryKey;autoIncrement:true"`
	Nama string
	Username string
	Password string
	Is_active bool
}