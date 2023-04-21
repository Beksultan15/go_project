package service

import (
	"github.com/Beksultan15/go_project"
	"github.com/Beksultan15/go_project/pkg/repository"
	"github.com/gin-gonic/gin"
)

type Authorization interface {
	CreateUser(user go_project.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	RefreshToken(username,password string) (string,error)
}

type Searhing interface {
	GetSearchingProduct(c *gin.Context) ([]go_project.Products,error)
}

type Service struct {
	Authorization
	Searhing
}

func NewService(repos *repository.Repository) *Service {
    return &Service{
		Authorization: newAuthService(repos.Authorization),
		Searhing: newSearchService(repos.Searching),
	}  
}