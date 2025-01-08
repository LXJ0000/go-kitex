package dal

import (
	"github.com/LXJ0000/go-kitex/app/user/biz/dal/mysql"
	"github.com/LXJ0000/go-kitex/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
