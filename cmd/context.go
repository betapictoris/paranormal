package main

type TokenContext struct {
	Username string
	Password string
}

func (token TokenContext) CreateToken() (string, error) {
	// TODO: Check if username and password are good.

	return GenerateRandomString(32)

	// TODO: Store the created token.
}
