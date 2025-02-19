package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TokenDetails struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

func Token(c *gin.Context) {
	guid := c.Query("guid")
	ip := c.ClientIP()
	fmt.Println("ip: ", ip)
	tokens, err := createToken(guid, ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	// _, err := uuid.Parse(guid)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid GUID format"})
	// 	return
	// }
	c.JSON(http.StatusOK, tokens)
}

func createToken(guid, ip string) (*TokenDetails, error) {
	accessTokenExpiration := time.Now().Add(30 * time.Minute).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": guid,
		"ip":      ip,
		"exp":     accessTokenExpiration,
	})

	accessTokenString, err := accessToken.SignedString([]byte("your_secret_key"))
	if err != nil {
		return nil, err
	}

	refreshToken := uuid.New().String()

	if err := saveRefreshToken(guid, refreshToken); err != nil {
		return nil, err
	}

	return &TokenDetails{
		AccessToken:  accessTokenString,
		RefreshToken: refreshToken,
		ExpiresAt:    accessTokenExpiration,
	}, nil
}

func saveRefreshToken(userID, refreshToken string) error {
	// hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	// if err != nil {
	// 	return err
	// }

	// db, err := db.getDB()

	// _, err = db.Exec(`INSERT INTO refresh_tokens (user_id, refresh_token_hash)
	//     VALUES ($1, $2)
	//     ON CONFLICT (user_id) DO UPDATE
	//     SET refresh_token_hash = $2, created_at = CURRENT_TIMESTAMP`, userID, string(hashedToken))

	// if err != nil {
	// 	log.Println("Error saving refresh token:", err)
	// 	return err
	// }
	return nil
}
