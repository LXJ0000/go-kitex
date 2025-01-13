// Code generated by hertz generator.

package main

import (
	"context"
	"os"
	"time"

	"github.com/LXJ0000/go-kitex/app/gateway/biz/router"
	"github.com/LXJ0000/go-kitex/app/gateway/conf"
	"github.com/LXJ0000/go-kitex/app/gateway/infra/rpc"
	"github.com/LXJ0000/go-kitex/common/mtl"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"github.com/joho/godotenv"
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
		hlog.Fatal("Error loading .env file")
	}
	registryAddr = os.Getenv("ETCD_ADDR")
	registerUserName = os.Getenv("ETCD_USERNAME")
	registerPassword = os.Getenv("ETCD_PASSWORD")

	serviceName = conf.GetConf().Hertz.Service
	metricsPort = conf.GetConf().Hertz.MetricsPort
}

func main() {
	// load .env
	initEnv()

	// mtl init
	r, info := mtl.InitMetric(serviceName, registryAddr, registerUserName, registerPassword, metricsPort)
	defer r.Deregister(info)

	p := mtl.InitTracing(serviceName)
	defer p.Shutdown(context.Background())

	// init rpc client
	rpc.Init()

	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)

	// add a ping route to test
	h.GET("/test", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"message": "ok"})
	})

	router.GeneratedRegister(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}
