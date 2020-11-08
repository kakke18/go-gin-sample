package handler

import (
	"fmt"
	"net/http"
)

type App struct{}

type NilArg struct{}

func (root *App) GetUsers(r *http.Request, arg *NilArg, reply *GetUsersReply) error {
	name := "GetUsers"
	fmt.Printf("[START] >> %s\n", name)
	defer fmt.Printf("[END] >> %s\n", name)
	return getUsers(r, arg, reply)
}
