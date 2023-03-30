package models

type User struct {
	ID          int64  `json:"id""`
	FullName    string `json:"full_name""`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	SocialID    string `json:"social_id"`
	Provider    string `json:"provider"`
	Role        int64  `json:"role"`
	OldPassword string `json:"old_password,omitempty"`
}
