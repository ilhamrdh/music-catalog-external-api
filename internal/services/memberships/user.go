package memberships

import (
	"errors"

	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	"github.com/ilhamrdh/music-catalog-external-api/pkg/jwt"
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

func (s *service) SignIn(request memberships.SignInRequest) (string, error) {
	user, err := s.repository.GetUser(request.Email, "", 0)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error().Err(err).Msg("error get user from database")
		return "", err
	}
	if user == nil {
		return "", errors.New("email or password invalid")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		log.Error().Err(err).Msg("email or password invalid")
		return "", errors.New("email or password invalid")
	}

	accessToken, err := jwt.GenerateToken(uint64(user.ID), user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		log.Error().Err(err).Msg("failed to create JWT token")
		return "", err
	}

	return accessToken, nil
}
