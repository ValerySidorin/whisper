package dto

type User struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
	Email     string `json:"email"`
}
