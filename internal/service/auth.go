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
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
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

// func saveRefreshToken(userID, refreshToken string) error {
// 	// hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// db, err := db.getDB()

// 	// _, err = db.Exec(`INSERT INTO refresh_tokens (user_id, refresh_token_hash)
// 	//     VALUES ($1, $2)
// 	//     ON CONFLICT (user_id) DO UPDATE
// 	//     SET refresh_token_hash = $2, created_at = CURRENT_TIMESTAMP`, userID, string(hashedToken))

// 	// if err != nil {
// 	// 	log.Println("Error saving refresh token:", err)
// 	// 	return err
// 	// }
// 	return nil
// }
