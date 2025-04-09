package models

type Animal interface {
	fala() string
	quantidadePatas() int
}

type User struct {
	Username string
	Password string
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

func (u *User) GetPassworld() string {
	return u.Password
}

func (u *User) SetPassword(passworld string) {
	u.Password = passworld
}
