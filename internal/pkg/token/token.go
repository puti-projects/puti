package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/puti-projects/puti/internal/pkg/config"
)

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
	// ErrMissingToken means the `token` params was empty.
	ErrMissingToken = errors.New("missing token")
)

// Context is the context of the JSON view token.
type Context struct {
	ID       uint64
	Username string
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	// Parse the token.
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	// Parse error.
	if err != nil {
		return ctx, err

		// Read the token if it's valid.
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.Username = claims["username"].(string)
		return ctx, nil

		// Other errors.
	} else {
		return ctx, err
	}
}

// ParseToken directly parse the token
func ParseToken(t string) (*Context, error) {
	// Load the jwt secret from config
	secret := config.Safety.JwtSecret

	if len(t) == 0 {
		return &Context{}, ErrMissingToken
	}

	return Parse(t, secret)
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	// Load the jwt secret from config
	secret := config.Safety.JwtSecret

	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	// Parse the header to get the token part.
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// Sign signs the context with the specified secret.
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	// Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = config.Safety.JwtSecret
	}
	// The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"username": c.Username,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	// Sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))

	return
}
