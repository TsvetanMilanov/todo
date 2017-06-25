package services

import (
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/TsvetanMilanov/todo/server/pkg/constants"
	"github.com/TsvetanMilanov/todo/server/pkg/db/models"
	"github.com/TsvetanMilanov/todo/server/pkg/types"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/golang/glog"
)

// TokensService token related methods.
type TokensService struct {
	DbService   types.IDbService   `inject:"dbService"`
	DateService types.IDateService `inject:"dateService"`
	Helpers     types.IHelpers     `inject:"helpers"`
}

// CreateToken creates a token for the user.
func (service *TokensService) CreateToken(username, tokenType string) (*models.Token, error) {
	now := service.DateService.Now()
	exp := now.Add(time.Second * time.Duration(constants.TokenExpirationTime))
	token := &models.Token{
		Nbf:       now.Unix(),
		Exp:       exp.Unix(),
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

	err = service.RemovExpiredTokens(username)
	if err != nil {
		glog.Infof("Remove expired tokens for user %s error: %s", username, err.Error())
	}

	tokens := service.DbService.GetCollection(constants.TokensCollectionName)
	err = tokens.Insert(token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

// ParseToken returns token object from the token string.s
func (service *TokensService) ParseToken(token string) (*models.SimpleTokenInfo, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return service.getTokenSigningSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token signature")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token signature")
	}

	// We need this conversion because for some reaseon the nbf and exp
	// are float64 (e.g. exp:1.499027492e+09 nbf:1.498422692e+09).
	nbf, err := service.Helpers.Float64ToInt64(claims["nbf"].(float64))
	if err != nil {
		return nil, err
	}

	exp, err := service.Helpers.Float64ToInt64(claims["exp"].(float64))
	if err != nil {
		return nil, err
	}

	scopesObject := claims["scopes"].([]interface{})
	var scopes []string

	for i, v := range scopesObject {
		scopes[i] = v.(string)
	}

	result := &models.SimpleTokenInfo{
		AccessToken: token,
		Nbf:         nbf,
		Exp:         exp,
		Scopes:      scopes,
		TokenType:   claims["token_type"].(string),
		Username:    claims["username"].(string),
	}

	return result, nil
}

// IsTokenExpired checks if the token has expired.
func (service *TokensService) IsTokenExpired(token string) bool {
	tokenInfo, err := service.ParseToken(token)
	if err != nil {
		return false
	}

	expDate := time.Unix(tokenInfo.Exp, 0)

	return expDate.Before(service.DateService.Now())
}

// RemovExpiredTokens removes all expired tokens of the provided user.
func (service *TokensService) RemovExpiredTokens(username string) error {
	tokens := service.DbService.GetCollection(constants.TokensCollectionName)
	return tokens.Remove(bson.M{"username": username, "exp": bson.M{"$lte": service.DateService.Now().Unix()}})
}

func (service *TokensService) generateJwtToken(tokenInfo *models.Token) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":        tokenInfo.Nbf,
		"exp":        tokenInfo.Exp,
		"username":   tokenInfo.Username,
		"token_type": tokenInfo.TokenType,
		"scopes":     tokenInfo.Scopes,
	})

	secret := service.getTokenSigningSecret()
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return "", nil
	}

	return signedToken, nil
}

func (service *TokensService) getTokenSigningSecret() []byte {
	secret := service.Helpers.GetEnvVariableOrDefault(constants.JwtSecret, constants.DefaultJwtSecret)
	return []byte(secret)
}
