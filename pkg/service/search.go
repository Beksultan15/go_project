package service

import (
	"github.com/Beksultan15/go_project"
	"github.com/Beksultan15/go_project/pkg/repository"
	"github.com/gin-gonic/gin"
)

type SearchService struct {
	repo repository.Searching
}

func newSearchService(repo repository.Searching) *SearchService{
	return &SearchService{repo:repo}
}

func (a *SearchService) GetSearchingProduct(c *gin.Context) ([]go_project.Products,error) {
	return a.repo.GetSearchingProduct(c)
	
}
