package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/ShekleinAleksey/jwt-auth/internal/entity"
	"github.com/ShekleinAleksey/jwt-auth/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

const (
	salt       = "hjqrhjqw124617ajfhajs"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL   = 12 * time.Hour
	refreshTTL = 30 * 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	userID, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (s *AuthService) CreateToken(email, password string) (string, string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", "", errors.New("user not found")
	}

	accessToken, refreshToken, err := s.GenerateToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) GenerateToken(user entity.User) (string, string, error) {

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	fmt.Println("accessToken: ", accessToken)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	accessTokenString, err := accessToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}

	err = s.repo.SaveRefreshToken(user.ID, refreshTokenString, refreshTTL)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GetUsers() ([]entity.User, error) {
	return s.repo.GetUsers()
}

func (s *AuthService) FindRefreshToken(userID int) (string, error) {
	refreshToken, err := s.repo.FindRefreshToken(userID)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *AuthService) FindUser(userID int) (entity.User, error) {
	user, err := s.repo.FindUserByID(userID)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (s *AuthService) RefreshToken(refreshToken string) (string, string, error) {
	userID, err := s.ParseToken(refreshToken)
	if err != nil {
		return "", "", err
	}

	user, err := s.FindUser(userID)
	if err != nil {
		return "", "", err
	}

	savedRefreshToken, err := s.FindRefreshToken(userID)
	if err != nil {
		return "", "", err
	}

	if savedRefreshToken != refreshToken {
		return "", "", errors.New("savedRefreshToken != refreshToken")
	}

	accessToken, newRefreshToken, err := s.GenerateToken(user)
	if err != nil {
		return "", "", err
	}

	return accessToken, newRefreshToken, nil
}

// func createToken(guid, ip string) (*TokenDetails, error) {
// 	accessTokenExpiration := time.Now().Add(30 * time.Minute).Unix()
// 	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"user_id": guid,
// 		"ip":      ip,
// 		"exp":     accessTokenExpiration,
// 	})

// 	accessTokenString, err := accessToken.SignedString([]byte("your_secret_key"))
// 	if err != nil {
// 		return nil, err
// 	}

// 	refreshToken := uuid.New().String()

// 	if err := saveRefreshToken(guid, refreshToken); err != nil {
// 		return nil, err
// 	}

// 	return &TokenDetails{
// 		AccessToken:  accessTokenString,
// 		RefreshToken: refreshToken,
// 		ExpiresAt:    accessTokenExpiration,
// 	}, nil
// }
