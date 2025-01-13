package mtl

import (
	"net"
	"net/http"

	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Registry *prometheus.Registry // 注册中心 用来注册指标

func InitMetric(serviceName, registryAddr, registryUserName, registryPassword, metricsPort string) (registry.Registry, *registry.Info) {
	Registry = prometheus.NewRegistry()
	Registry.MustRegister(collectors.NewGoCollector())
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	r, err := etcd.NewEtcdRegistryWithAuth([]string{registryAddr}, registryUserName, registryPassword)
	// r, err := consul.NewConsulRegister(registryAddr)
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "localhost"+metricsPort)
	if err != nil {
		panic(err)
	}
	registryInfo := &registry.Info{ // 注册的信息
		ServiceName: "prometheus", // 服务名 把指标服务统一注册到名为 prometheus 的服务中心
		Addr:        addr,
		Weight:      1,
		Tags: map[string]string{
			"service": serviceName,
		},
	}
	if err := r.Register(registryInfo); err != nil {
		panic(err)
	}
	server.RegisterShutdownHook(func() {
		if err := r.Deregister(registryInfo); err != nil {
			panic(err)
		}
	})

	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	go http.ListenAndServe(metricsPort, nil) // 启动一个 http 服务，用来被 prometheus 抓取指标

	return r, registryInfo
}
