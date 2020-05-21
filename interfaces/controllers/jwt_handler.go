package controllers

type JWTHandler interface {
	Verify(string) (string, error)
}
