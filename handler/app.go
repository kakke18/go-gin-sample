package handler

import (
	"fmt"
	"net/http"
)

type App struct{}

type NilArg struct{}

func (root *App) GetMembers(r *http.Request, arg *NilArg, reply *GetMembersReply) error {
	name := "GetMembers"
	fmt.Printf("[START] >> %s\n", name)
	defer fmt.Printf("[END] >> %s\n", name)
	return getMembers(r, arg, reply)
}
