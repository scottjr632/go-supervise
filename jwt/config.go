package jwt

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Config defaults
var (
	DefaultExpirationtime = time.Now().Add(8 * time.Hour)
	DefaultJWTKey         = []byte("my_secret_key") // only here for testing purposes
	DefaultSigningMethod  = jwt.SigningMethodHS256
	DefaultTokenName      = "id_token"

	// DefaultTokenExtractor returns the token from the cookie
	DefaultTokenExtractor = func(tokenName string, r *http.Request) (string, error) {
		cookie, err := r.Cookie(tokenName)
		if err != nil {
			return "", err
		}

		return cookie.Value, nil
	}
)

// Config ...
type Config struct {
	ExpirationTime time.Time
	JWTKey         []byte
	SigningMethod  *jwt.SigningMethodHMAC
	TokenName      string
	ProtectedPath  string

	TokenExtractor func(tokenName string, r *http.Request) (string, error)
}

// ApplyDefaults applys the default values to empty configs
func applyDefaults(config *Config) {
	if config.ExpirationTime.Day() == 1 {
		config.ExpirationTime = DefaultExpirationtime
	}
	if len(config.JWTKey) == 0 {
		config.JWTKey = DefaultJWTKey
	}
	if config.TokenExtractor == nil {
		config.TokenExtractor = DefaultTokenExtractor
	}
	if config.TokenName == "" {
		config.TokenName = DefaultTokenName
	}
	config.SigningMethod = DefaultSigningMethod
}
