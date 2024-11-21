package memberships

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedBy string `gorm:"not null"`
	UpdatedBy string `gorm:"not null"`
}

type (
	SignUpRequest struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	SignInRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)

type (
	SignInResponse struct {
		AccessToken string `json:"access_token"`
	}
)
