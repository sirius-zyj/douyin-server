package jwt

import (
	"douyin-server/config"
	"douyin-server/database/dao"
	"log"
	"strconv"
	"time"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(username string, delaytime int) string {
	time.Sleep(time.Second * time.Duration(delaytime))
	user, err := dao.GetUsersByUserName(username)
	if err != nil {
		log.Println("获取用户失败")
		return "用户获取失败"
	}
	token := NewToken(user)
	return token
}

func NewToken(u dao.Duser) string {
	expiresTime := time.Now().Unix() + int64(config.OneDayOfHours)
	// log.Printf("expiresTime: %v\n", expiresTime)
	id64 := u.ID
	// log.Printf("id: %v\n", strconv.FormatInt(id64, 10))
	claims := jwt.StandardClaims{
		Audience:  u.Name,
		ExpiresAt: expiresTime,
		Id:        strconv.FormatInt(id64, 10),
		IssuedAt:  time.Now().Unix(),
		Issuer:    "tiktok",
		NotBefore: time.Now().Unix(),
		Subject:   "token",
	}
	var jwtSecret = []byte(config.Secret)
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err == nil {
		// println("generate token success!\n")
		return token
	} else {
		log.Fatalln("generate token fail")
		return "fail"
	}
}

func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))

	log.Println("Result: " + sha)
	return sha
}

func GetUserIdByToken(tokenString string) int64 {
	// 解析 JWT token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		panic("Failed to parse token: " + err.Error())
	}

	// 验证 token 是否有效
	if !token.Valid {
		panic("Invalid token")
	}

	// 提取用户 ID
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		panic("Failed to extract claims")
	}
	userID := claims.Id
	id, _ := strconv.ParseInt(userID, 10, 64)
	return id
}
