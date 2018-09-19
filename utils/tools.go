package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/smtp"
	"regexp"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CheckEmail check email is right format
func CheckEmail(email string) bool {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegexp.MatchString(email)
}

// SendEmail send email to user
func SendEmail(email string, checkToken string) error {
	auth := smtp.PlainAuth("", "customer@chelecom.io", "Vhsej5UMWBcZQcwf", "smtp.exmail.qq.com")
	to := []string{email}
	nickname := "emailNickname"
	user := "customer@chelecom.io"
	subject := "Active Email"
	contentType := "Content-Type: text/html; charset=UTF-8"
	body := "<html><body><a href=http://192.168.1.215:3000/#/activeAccount?token=" + checkToken + ">Active http://localhost:8888/#/activeAccount?token=" + checkToken + "</a><body></html>"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + nickname +
		"<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)

	err := smtp.SendMail("smtp.exmail.qq.com:25", auth, user, to, msg)
	return err
}

// CryptoPassword crypto password
func CryptoPassword(data string) string {
	h := hmac.New(sha256.New, []byte("userPassword"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

// CreateTokenEndpoint create jwt-token
func CreateTokenEndpoint(ID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  ID,
		"exp": time.Now().Add(30 * time.Minute).Unix(),
	})
	tokenString, error := token.SignedString([]byte("secret"))
	return tokenString, error
}

// ProtectedEndpoint protect jwt-token
func ProtectedEndpoint(params string) (bool, int) {
	token, _ := jwt.Parse(params, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		idStr := int(claims["id"].(float64))
		return true, idStr
	} else {
		return false, 0
	}
}
