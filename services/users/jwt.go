package users

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kacpertarka/restaurant/utils"
)

type JWToken struct {
	secretKey       string
	accesTokenExp   int64 // in seconds??
	refreshTokenExp int64
}

func NewJWT(secretKey string, accesTokenExp, refreshTokenExp int64) *JWToken {
	return &JWToken{
		secretKey:       utils.GetEnvVariable("JWT_SECRET_KEY", "very_secre_key"),
		accesTokenExp:   int64(utils.GetEnvVariableAsInt("JWT_ACCESS_EXP", 60*10)),
		refreshTokenExp: int64(utils.GetEnvVariableAsInt("JWT_REFRESH_EXP", 60*10*6)),
	}
}

func (j *JWToken) GenerateToken(userID string) (*TokenResponse, error) {
	accesToken, err := j.generateToken(userID, j.accesTokenExp)
	refreshToken, err := j.generateToken(userID, j.refreshTokenExp)
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		TokenType:    "Bearer",
		AccessToken:  accesToken,
		RefreshToken: refreshToken}, nil
}

func (j *JWToken) ValidateToken(tokenString string) error {
	// now return nil if token is valid - in the future return data from token? email?
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func (j *JWToken) generateToken(userID string, expTime int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Second * time.Duration(j.accesTokenExp)).Unix(),
	})
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
