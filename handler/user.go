package handler

import (
	"go-gin-sample/dao"
	"go-gin-sample/models"
	"net/http"
)

// GetUsers
type GetUsersArg struct{}

type GetUsersReply struct {
	Users []models.User `json:"users"`
}

func getUsers(r *http.Request, args *GetUsersArg, reply *GetUsersReply) error {
	userDao := dao.NewUserJsonDao()
	results, err := userDao.List()
	if err != nil {
		return err
	}

	reply.Users = *results
	return nil
}

// GetUser
type GetUserArg struct {
	Id string `json:"id"`
}

type GetUserReply struct {
	User models.User `json:"user"`
}

func getUser(r *http.Request, args *GetUserArg, reply *GetUserReply) error {
	userDao := dao.NewUserJsonDao()
	results, err := userDao.Get(args.Id)
	if err != nil {
		return err
	}

	reply.User = *results
	return nil
}
