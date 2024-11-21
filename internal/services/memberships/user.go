package memberships

import (
	"errors"

	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (s *service) SignUp(request memberships.SignUpRequest) error {
	existingUser, err := s.repository.GetUser(request.Email, request.Username, 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get user from database")
		return err
	}

	if existingUser != nil {
		return errors.New("email or username exists")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("error hash password")
		return err
	}

	model := memberships.User{
		Email:     request.Email,
		Username:  request.Username,
		Password:  string(hash),
		CreatedBy: request.Username,
		UpdatedBy: request.Username,
	}
	return s.repository.CreateUser(model)
}
