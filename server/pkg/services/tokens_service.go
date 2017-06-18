package services

import (
	"time"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// TokensService token related methods.
type TokensService struct {
	DbService types.IDbService `inject:"dbService"`
	Helpers   types.IHelpers   `inject:"helpers"`
}

// CreateToken creates a token for the user.
func (service *TokensService) CreateToken(username, tokenType string) (*models.Token, error) {
	nbf := time.Now().Add(time.Second * time.Duration(constants.TokenExpirationTime))
	token := models.Token{
		ID:        uuid.New().String(),
		Nbf:       nbf.Unix(),
		ExpiresIn: constants.TokenExpirationTime,
		Scopes:    []string{},
		TokenType: tokenType,
		Username:  username,
	}

	accessToken, err := service.generateJwtToken(token)
	if err != nil {
		return nil, err
	}

	token.AccessToken = accessToken

	tokens := service.DbService.GetCollection(constants.TokensCollectionName)
	err = tokens.Insert(token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (service *TokensService) generateJwtToken(tokenInfo models.Token) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":        tokenInfo.Nbf,
		"username":   tokenInfo.Username,
		"token_type": tokenInfo.TokenType,
		"scopes":     tokenInfo.Scopes,
	})

	secret := service.Helpers.GetEnvVariableOrDefault(constants.JwtSecret, constants.DefaultJwtSecret)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", nil
	}

	return signedToken, nil
}
