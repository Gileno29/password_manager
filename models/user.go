package models

func NewUser(username string, password string, token string, tipo *conta) *user {
	return &user{
		Username: username,
		Password: password,
		Token:    token,
		Tipo:     tipo,
	}
}

type user struct {
	Username string
	Password string
	Token    string
	Tipo     *conta
}

func (u *user) GetUsername() string {
	return u.Username
}

func (u *user) SetUsername(username string) {
	u.Username = username
}

func (u *user) GetPassworld() string {
	return u.Password
}

func (u *user) SetPassword(passworld string) {
	u.Password = passworld
}
