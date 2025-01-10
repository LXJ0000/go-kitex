package serversuite

import (
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type CommonServerSuite struct {
	CurrentServiceName string
	RegistryAddr       string
	RegistryAuthUserName string
	RegistryAuthPassword string
}

func (s CommonServerSuite) Options() []server.Option {
	opts := []server.Option{
		// service info
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
	}

	// etcd
	r, err := etcd.NewEtcdRegistryWithAuth([]string{s.RegistryAddr}, s.RegistryAuthUserName, s.RegistryAuthPassword)
	if err != nil {
		klog.Fatalf("new etcd registry failed: %v", err)
	}
	opts = append(opts, server.WithRegistry(r))

	return opts
}
