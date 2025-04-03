package usecases

import (
	"auth/common/config"
	"auth/modules/user/data/model"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	GenerateToken(user *model.User) (string, error)
}

type jwtServiceImpl struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTService(config config.Config) JWTService {
	cfg := config
	return &jwtServiceImpl{
		secretKey:     cfg.JwtSecret,
		tokenDuration: time.Duration(cfg.TokenDuration) * time.Hour,
	}
}

func (s *jwtServiceImpl) GenerateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":      user.ID.Hex(),
		"community_id": user.CommunityID.Hex(),
		"role_id":      user.RoleID.Hex(),
		"status":       user.Status,
		"exp":          time.Now().Add(s.tokenDuration).Unix(),
		"iat":          time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
	
}