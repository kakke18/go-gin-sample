package handler

import (
	"go-gin-sample/dao"
	"go-gin-sample/models"
	"net/http"
)

type GetMembersArg struct{}

type GetMembersReply struct {
	Members []models.Member `json:"members"`
}

func getMembers(r *http.Request, args *NilArg, reply *GetMembersReply) error {
	memberDao := dao.NewMemberJsonDao()
	results, err := memberDao.Get()
	if err != nil {
		return err
	}

	reply.Members = *results
	return nil
}
