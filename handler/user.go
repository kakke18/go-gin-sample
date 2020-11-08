package handler

import (
	"go-gin-sample/dao"
	"go-gin-sample/models"
	"net/http"
)

type GetUsersArg struct{}

type GetUsersReply struct {
	Users []models.User `json:"users"`
}

func getUsers(r *http.Request, args *NilArg, reply *GetUsersReply) error {
	userDao := dao.NewUserJsonDao()
	results, err := userDao.Get()
	if err != nil {
		return err
	}

	reply.Users = *results
	return nil
}
