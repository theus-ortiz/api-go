package responses

type AuthResponse struct {
    Token string `json:"token"`
    Message string `json:"message"`
}