package token

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/lopo1/mall/auth/utils"

	"github.com/dgrijalva/jwt-go"
)
var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)
// JWTTokenVerifier verifies jwt access tokens.
type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

// Verify verifies a token and returns account id.
func (v *JWTTokenVerifier) Verify(token string) (*utils.CustomClaims, error) {
	t, err := jwt.ParseWithClaims(token, &utils.CustomClaims{},
		func(*jwt.Token) (interface{}, error) {
			return v.PublicKey, nil
		})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}


	if !t.Valid {
		return nil, fmt.Errorf("token not valid")
	}

	clm, ok := t.Claims.(*utils.CustomClaims)
	if !ok {
		return nil, fmt.Errorf("token claim is not StandardClaims")
	}

	if err := clm.Valid(); err != nil {
		return nil, fmt.Errorf("claim not valid: %v", err)
	}

	return clm, nil
}
