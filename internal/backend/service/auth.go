package service

// Token represents a JSON web token.
type Token struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
