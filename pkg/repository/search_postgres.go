package repository

import (
	"fmt"
	"github.com/Beksultan15/go_project"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/postgres"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


type SearchPostgres struct {
	dba *sqlx.DB
}


func NewSearchPostgres(dba *sqlx.DB) *SearchPostgres {
	return &SearchPostgres{dba:dba}
}


func (a *SearchPostgres) GetSearchingProduct(c *gin.Context) ([]go_project.Products, error) {
	dsn := "host=postgres-db user=docker password=docker dbname=test_db port=5432 sslmode=disable"
	
	db,err := gorm.Open(postgres.Open(dsn),&gorm.Config{})
	var products []go_project.Products
		
	if err != nil {
		panic("Could not connect to the database")
	}
	sql := "SELECT fp.id,fp.producttitle,fp.price,fp.quantity,fp.pddesc, pr.image FROM products AS fp join furni_productimages AS pr on fp.id = pr.product_id"
	
		if s := c.Query("s"); s != "" {
			sql = fmt.Sprintf("%s WHERE fp.producttitle LIKE '%%%s%%' LIMIT 1", sql, s)
		}
		db.Raw(sql).Scan(&products)

		return products,nil
}


