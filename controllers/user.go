package controllers

import (
	"context"
	"net/http"

	"github.com/aminGhafoory/webshop/models"
	"github.com/google/uuid"
)

type Users struct {
	UserService *models.UserService
}

func (u *Users) TestHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := u.UserService.DB.GetAuthor(context.Background(), uuid.MustParse("07071006-e7d8-4aaa-8e5b-c9a1f51c7987"))
	w.Write([]byte(user.Email))
}
