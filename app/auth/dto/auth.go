package dto

import "github.com/wisnunu254/api-auth-golang/app/auth/model"

type AuthDTOLogin struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	// Type  string `json:"type"`
}

type TokensDTO struct {
	AccessToken string `json:"access"`
}

func UserToDTO(user *model.User) AuthDTOLogin {
	return AuthDTOLogin{
		ID:    user.ID,
		Email: user.Email,
		// Type:  user.Type,
		// Add other fields as needed
	}
}

func TokenToDTO(token string) TokensDTO {
	return TokensDTO{
		AccessToken: token,
	}
}
