package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-sample/dao"
)

func main() {
	router := gin.Default()
	router.GET("/members", func(c *gin.Context){
		memberDao := dao.NewMemberJsonDao()
		results, err := memberDao.Get()

		var message interface{}
		if err != nil {
			message = err
		} else {
			message = results
		}

		c.JSON(200, gin.H{ "message": message})
	})
	_ = router.Run()
}
