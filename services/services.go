package services

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/sethvargo/go-password/password"
)

const (
	StatusTokenOK      = 0
	StatusTokenInvalid = 1
	StatusBadToken     = 2
	StatusTokenExpired = 3
	serverKey          = "AAAAn5-NFHs:APA91bFWjxqVscp0zON2v3UBGrdxwheQY610qwwZlJ8wPRBKHD-wXVr15C0ccefibA1_hbQ5oR0tajW9YBZe1a_PzLfsH41I2Or1DkWjfrCPxCWnvPzL-lzUpyGtzkdVLr33W4gwcWmM"
)

type UserRoleJwt struct {
	jwt.StandardClaims
	UserID int64 `json:"user_id"`
	RoleID int64 `json:"role_id"`
}

func (claims *UserRoleJwt) CreateToken(key string) (token string, expTime int64) {
	if claims.ExpiresAt == 0 {
		expTime = time.Now().UTC().AddDate(0, 6, 0).Unix()
		claims.ExpiresAt = expTime
	} else {
		expTime = claims.ExpiresAt
	}

	var err error
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = t.SignedString([]byte(key))
	if err != nil {
		return "", 0
	}

	return token, expTime
}

func (claims *UserRoleJwt) ParseToken(token string, key string) int {
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return StatusTokenInvalid
		}
		return StatusBadToken
	}

	if !t.Valid {
		return StatusTokenInvalid
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now().UTC()) < 0 {
		return StatusTokenExpired
	}

	return StatusTokenOK
}

func HashAndSalt(pwd string) (hash string, salt string) {
	var err error
	salt, err = password.Generate(64, 10, 10, false, true)
	if err != nil {
		return "", ""
	}

	salted := pwd + salt
	result := sha256.Sum256([]byte(salted))
	hash = fmt.Sprintf("%x", result)
	return hash, salt
}

func IsPasswordMatch(pwd string, pwdHash string, salt string) bool {
	salted := pwd + salt
	result := sha256.Sum256([]byte(salted))
	hash := fmt.Sprintf("%x", result)
	return hash == pwdHash
}

// func Notif(token, title, body string, order_id int64) {
// 	msg := &fcm.Message{
// 		To: token,
// 		Data: map[string]interface{}{
// 			"order_id": order_id,
// 		},
// 		Notification: &fcm.Notification{
// 			Title: title,
// 			Body:  body,
// 		},
// 		ContentAvailable: true,
// 	}

// 	// Create a FCM client to send the message.
// 	client, err := fcm.NewClient(serverKey)
// 	if err != nil {
// 		fmt.Println(err)

// 	}

// 	// Send the message and receive the response without retries.
// 	response, err := client.Send(msg)
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	fmt.Println(response)
// }
