package clientsuite

import (
	"log"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	etcd "github.com/kitex-contrib/registry-etcd"
)

type CommonClientSuite struct {
	CurrentServiceName   string
	RegistryAddr         string
	RegistryAuthUserName string
	RegistryAuthPassword string
}

func (s CommonClientSuite) Options() []client.Option {
	opts := []client.Option{
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: s.CurrentServiceName,
		}),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
	}

	// etcd
	r, err := etcd.NewEtcdResolverWithAuth([]string{s.RegistryAddr}, s.RegistryAuthUserName, s.RegistryAuthPassword)
	if err != nil {
		log.Fatal(err)
	}
	opts = append(opts, client.WithResolver(r))

	return opts
}
