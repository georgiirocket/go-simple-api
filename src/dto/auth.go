package dto

type AuthData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessExpire int64  `json:"accessExpire"`
	ExpiresIn    int64  `json:"expiresIn"`
}

type AuthPayload struct {
	TokenType string
	UserId    string
	Role      string
	Exp       int64
}

type AuthPayloadRefresh struct {
	TokenType       string
	UserId          string
	Role            string
	Exp             int64
	PartAccessToken string
}
