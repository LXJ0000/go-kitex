package main

import (
	"net"
	"os"
	"time"

	"github.com/LXJ0000/go-kitex/app/user/biz/dal"
	"github.com/LXJ0000/go-kitex/app/user/conf"
	"github.com/LXJ0000/go-kitex/common/serversuite"
	"github.com/LXJ0000/go-kitex/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	opts := kitexInit()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// load .env
	if err := godotenv.Load(); err != nil {
		klog.Fatal("Error loading .env file")
	}

	// init dal
	dal.Init()

	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(
		opts,
		server.WithServiceAddr(addr),
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName:   conf.GetConf().Kitex.Service,
			RegistryAddr:         os.Getenv("ETCD_ADDR"),
			RegistryAuthUserName: os.Getenv("ETCD_USERNAME"),
			RegistryAuthPassword: os.Getenv("ETCD_PASSWORD"),
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
