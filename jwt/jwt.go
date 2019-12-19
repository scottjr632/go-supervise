package jwt

import "github.com/dgrijalva/jwt-go"

type key int

const (
	keyPrincipalID key = iota
)

// JWT ...
type JWT struct {
	Config

	Key key
}

// Claims ...
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// NewJWTFromConfig returns a new JWT from config
func NewJWTFromConfig(config Config) *JWT {
	applyDefaults(&config)
	j := &JWT{config, keyPrincipalID}

	return j
}

// NewToken retuns a new JWT from the config
func (j *JWT) NewToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: j.ExpirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(j.SigningMethod, claims)
	return token.SignedString(j.JWTKey)
}

// ValidateToken validates the token and returns the claims
func (j *JWT) ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return j.JWTKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, ErrTokenNotValid
	}

	return claims, nil
}
