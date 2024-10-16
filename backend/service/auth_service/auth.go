package auth_services

import (
	"discord-clone/models"
	"discord-clone/pkg/util"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Check ...
func (a *Auth) Check() (bool, error) {
	user, err := models.GetUserByEmail(a.Email)
	if err != nil {
		return false, err
	}
	check := util.ComparePassword(user.PasswordHash, a.Password)
	return check, nil
}

// Get ...
func (a *Auth) GetUserPublic() (*models.UserPublic, error) {
	var User *models.UserPublic
	user, err := models.GetUserByEmail(a.Email)
	if err != nil {
		return nil, err
	}
	User = &models.UserPublic{
		ID:                user.ID,
		Username:          user.Username,
		Email:             user.Email,
		ProfilePictureUrl: user.ProfilePictureUrl,
	}
	return User, nil
}
