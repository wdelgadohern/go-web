package auth

type AuthTokenMock struct {
	FuncAuth func(token string) (err error)
}

func NewAuthTokenMock() *AuthTokenMock {
	return &AuthTokenMock{}
}

func (a *AuthTokenMock) Auth(token string) (err error) {
	return a.FuncAuth(token)
}
