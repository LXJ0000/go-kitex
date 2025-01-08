package service

import (
	"context"

	common "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/common"
	ping "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/gateway/ping"
	"github.com/cloudwego/hertz/pkg/app"
)

type PingService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewPingService(Context context.Context, RequestContext *app.RequestContext) *PingService {
	return &PingService{RequestContext: RequestContext, Context: Context}
}

func (h *PingService) Run(req *common.Empty) (resp *ping.PingResponse, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return &ping.PingResponse{Message: "pong"}, nil
}
