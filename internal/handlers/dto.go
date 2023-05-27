package handlers

import "university/internal/entity"

type ResponseSignIn struct {
	User  *entity.User
	Token string
}
