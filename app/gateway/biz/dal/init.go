package dal

import (
	"github.com/LXJ0000/go-kitex/app/gateway/biz/dal/mysql"
	"github.com/LXJ0000/go-kitex/app/gateway/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
