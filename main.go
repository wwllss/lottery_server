package main

import (
	"github.com/wwllss/zop"
	"lottery_server/config"
	"lottery_server/dao"
	"lottery_server/lottery"
)

func main() {
	db, _ := dao.GetDB().DB()
	defer func() {
		_ = db.Close()
	}()
	z := zop.New()
	z.Use(RtTime)
	z.Use(Recovery)
	lottery.Register(z)
	c := config.GetConfig()
	if err := z.Run(":" + c.Zop.Port); err != nil {
		panic(err)
	}
}
