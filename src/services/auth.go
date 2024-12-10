package services

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go-simple/src/dto"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

const (
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

func CreateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

func VerifyPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false
	}

	return true
}

// CreateAuthData . access 3h refresh 7days
func CreateAuthData(userId string, role string) (*dto.AuthData, error) {
	accessExpire := time.Now().Add(time.Duration(3) * time.Hour).Unix()
	expiresIn := time.Now().Add(time.Duration(168) * time.Hour).Unix()
	authData := dto.AuthData{AccessExpire: accessExpire, ExpiresIn: expiresIn}

	accessToken, errAccessToken := createAccessToken(&dto.AuthPayload{UserId: userId, Role: role, TokenType: AccessTokenType, Exp: accessExpire})

	if errAccessToken != nil {
		return nil, errAccessToken
	}

	authData.AccessToken = accessToken

	partAccessToken := accessToken[len(accessToken)-8:]

	refreshToken, errRefreshToken := createRefreshToken(&dto.AuthPayloadRefresh{UserId: userId, Role: role, TokenType: RefreshTokenType, Exp: expiresIn, PartAccessToken: partAccessToken})

	if errRefreshToken != nil {
		return nil, errRefreshToken
	}

	authData.RefreshToken = refreshToken

	return &authData, nil
}

func VerifyToken(accessToken string, tokenType string) (*dto.AuthPayload, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims["tokenType"].(string) != tokenType {
		return nil, errors.New("invalid token type")
	}

	payload := &dto.AuthPayload{
		UserId:    claims["userId"].(string),
		Role:      claims["role"].(string),
		TokenType: claims["tokenType"].(string),
		Exp:       int64(claims["exp"].(float64)),
	}

	return payload, nil
}

func VerifyRefreshToken(refreshToken string, accessToken string) (*dto.AuthPayloadRefresh, error) {
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	if claims["tokenType"].(string) != RefreshTokenType {
		return nil, errors.New("invalid token type")
	}

	payload := &dto.AuthPayloadRefresh{
		UserId:          claims["userId"].(string),
		Role:            claims["role"].(string),
		TokenType:       claims["tokenType"].(string),
		PartAccessToken: claims["partAccessToken"].(string),
		Exp:             int64(claims["exp"].(float64)),
	}

	partAccessToken := accessToken[len(accessToken)-8:]

	if partAccessToken != payload.PartAccessToken {
		return nil, errors.New("invalid pair")
	}

	return payload, nil
}

func createAccessToken(auth *dto.AuthPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId":    auth.UserId,
			"role":      auth.Role,
			"tokenType": auth.TokenType,
			"exp":       auth.Exp,
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func createRefreshToken(auth *dto.AuthPayloadRefresh) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId":          auth.UserId,
			"role":            auth.Role,
			"tokenType":       auth.TokenType,
			"partAccessToken": auth.PartAccessToken,
			"exp":             auth.Exp,
		})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
