package handler

import (
	"fmt"
	"net/http"
)

type App struct{}

func (root *App) GetUsers(r *http.Request, arg *GetUsersArg, reply *GetUsersReply) error {
	name := "GetUsers"
	fmt.Printf("[START] >> %s\n", name)
	defer fmt.Printf("[END] >> %s\n", name)
	return getUsers(r, arg, reply)
}

func (root *App) GetUser(r *http.Request, arg *GetUserArg, reply *GetUserReply) error {
	name := "GetUser"
	fmt.Printf("[START] >> %s\n", name)
	defer fmt.Printf("[END] >> %s\n", name)
	return getUser(r, arg, reply)
}
