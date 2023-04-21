package repository

import (
	"fmt"
	"time"

	"github.com/Beksultan15/go_project"
	"github.com/jmoiron/sqlx"
)


type AuthPostgres struct {
	db *sqlx.DB
}


func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db:db}
}

func (a *AuthPostgres) CreateUser(user go_project.User) (int, error) {
	var id int
	
	query := fmt.Sprintf("INSERT INTO %s (first_name,last_name,username,email,password,is_staff,is_superuser,is_active,date_joined) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id", usertable)
	row := a.db.QueryRow(query,user.First_name,user.Last_name,user.First_name,user.Email,user.Password,true,false,true,time.Now())
	
	if err:= row.Scan(&id); err!= nil {
		return 0,err
	}
	return id,nil
}


func (a *AuthPostgres) GetUser(username,password string) (go_project.User, error) {
	var user go_project.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usertable)
	err := a.db.Get(&user,query,username,password)
	return user,err
}