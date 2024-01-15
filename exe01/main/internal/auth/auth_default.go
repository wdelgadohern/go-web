package auth

type AuthBasic struct {
	Token string
}

func NewAuthBasic(token string) *AuthBasic {
	return &AuthBasic{
		Token: token,
	}
}

func (a *AuthBasic) Auth(token string) (err error) {
	if a.Token != token {
		return ErrAuthTokenInternal
	}
	return nil
}
