package main

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/LXJ0000/go-kitex/app/user/biz/dal"
	"github.com/LXJ0000/go-kitex/app/user/conf"
	"github.com/LXJ0000/go-kitex/common/mtl"
	"github.com/LXJ0000/go-kitex/common/serversuite"
	"github.com/LXJ0000/go-kitex/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	serviceName      string
	metricsPort      string
	registryAddr     string
	registerUserName string
	registerPassword string
)

func initEnv() {
	// load .env
	if err := godotenv.Load(); err != nil {
		klog.Fatal("Error loading .env file")
	}
	registryAddr = os.Getenv("ETCD_ADDR")
	registerUserName = os.Getenv("ETCD_USERNAME")
	registerPassword = os.Getenv("ETCD_PASSWORD")

	serviceName = conf.GetConf().Kitex.Service
	metricsPort = conf.GetConf().Kitex.MetricsPort
}

func main() {
	// load .env
	initEnv()

	// tracing init
	p := mtl.InitTracing(serviceName)
	defer p.Shutdown(context.Background())

	// mtl init
	r, info := mtl.InitMetric(serviceName, registryAddr, registerUserName, registerPassword, metricsPort)
	defer r.Deregister(info)

	// dal init
	dal.Init()

	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(
		opts,
		server.WithServiceAddr(addr),
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName:   serviceName,
			RegistryAddr:         registryAddr,
			RegistryAuthUserName: registerUserName,
			RegistryAuthPassword: registerPassword,
		}),
	)

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
