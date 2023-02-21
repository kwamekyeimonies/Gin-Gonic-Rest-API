package response

type UserResponse struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}


type LoginRespone struct{
	TokenType string `json:"token_type"`
	Token string `json:"token"`
}