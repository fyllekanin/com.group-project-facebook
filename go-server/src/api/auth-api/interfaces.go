package auth_api

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterPayload struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Repassword string `json:"repassword"`
}

type LoginResponse struct {
	Username     string `json:"username"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
