package models

// Token access token information.
// Don't use SimpleTokenInfo because the mongo driver gets confused.
type Token struct {
	ID          string   `json:"id" bson:"_id"`
	AccessToken string   `json:"access_token" bson:"access_token"`
	TokenType   string   `json:"token_type" bson:"token_type"`
	ExpiresIn   int      `json:"expires_in" bson:"expires_in"`
	Nbf         int64    `json:"nbf" bson:"nbf"`
	Exp         int64    `json:"exp" bson:"exp"`
	Username    string   `json:"username" bson:"username"`
	Scopes      []string `json:"scopes" bson:"scopes"`
}

// SimpleTokenInfo simple access token information.
type SimpleTokenInfo struct {
	AccessToken string   `json:"access_token" bson:"access_token"`
	TokenType   string   `json:"token_type" bson:"token_type"`
	Nbf         int64    `json:"nbf" bson:"nbf"`
	Exp         int64    `json:"exp" bson:"exp"`
	Username    string   `json:"username" bson:"username"`
	Scopes      []string `json:"scopes" bson:"scopes"`
}
