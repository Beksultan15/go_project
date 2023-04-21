package service

import (
	"github.com/Beksultan15/go_project"
	"github.com/Beksultan15/go_project/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
	"fmt"
	"errors"
	"crypto/sha1"
	"github.com/google/uuid"
)

const (salt = "at%^o*4y5d=&6rg58bvb"
	   tokenTTL = time.Hour * 1
	   signingKey = "django-insecure-@l(*i3aqztc#397#p%^8babo$$o1$9ct%^o*4y5d=&6rg58bvb"
	)


type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo:repo}
}

func (a *AuthService) CreateUser(user go_project.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)	 
    return a.repo.CreateUser(user)
}

func (a *AuthService) GenerateToken(username,password string) (string, error) {
	user, err := a.repo.GetUser(username,generatePasswordHash(password))
	if err!= nil {
        return "", err
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token_type":"access",
        "user_id": user.Id,
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(tokenTTL).Unix(),
        "jti": uuid.New().String(),
    })
	
	tokenString, err := token.SignedString([]byte(signingKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}


func (a *AuthService) RefreshToken(username,password string) (string, error) {
	user, err := a.repo.GetUser(username,generatePasswordHash(password))

	if err!= nil {
        return "", err
    }
	
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"token_type":"refresh",
        "user_id": user.Id,
        "iat": time.Now().Unix(),
        "exp": time.Now().Add(7 * 24* time.Hour).Unix(),
        "jti": uuid.New().String(),
    })
	
	tokenString, err := token.SignedString([]byte(signingKey))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x",hash.Sum([]byte(salt)))
}

func(s*AuthService) ParseToken(accessToken string) (int, error) {
	token,err:= jwt.ParseWithClaims(accessToken,&tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
	  if _, ok := token.Method.(*jwt.SigningMethodHMAC);!ok {	
		return nil, errors.New("invalid signing method")
	}	
	return []byte(signingKey),nil
	})
	if err!= nil {
			return 0, err
		}

		claims,ok :=token.Claims.(*tokenClaims)
		if !ok {
			return 0, errors.New("token claims are not of type ")
        }
		return claims.UserId, nil
		}
	



