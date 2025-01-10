package rpc

import (
	"os"
	"sync"

	"github.com/LXJ0000/go-kitex/app/gateway/conf"
	"github.com/LXJ0000/go-kitex/app/gateway/utils"
	"github.com/LXJ0000/go-kitex/common/clientsuite"
	"github.com/LXJ0000/go-kitex/rpc_gen/kitex_gen/user/userservice"

	"github.com/cloudwego/kitex/client"
)

var (
	UserClient userservice.Client

	once sync.Once
)

func Init() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	c, err := userservice.NewClient("user", client.WithSuite(clientsuite.CommonClientSuite{
		CurrentServiceName:   conf.GetConf().Hertz.Address,
		RegistryAddr:         os.Getenv("ETCD_ADDR"),
		RegistryAuthUserName: os.Getenv("ETCD_USERNAME"),
		RegistryAuthPassword: os.Getenv("ETCD_PASSWORD"),
	}))
	utils.MustHandleError(err)

	UserClient = c
}
