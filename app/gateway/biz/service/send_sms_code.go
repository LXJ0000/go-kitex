package service

import (
	"context"

	common "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/common"
	auth2 "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/gateway/auth2"
	"github.com/cloudwego/hertz/pkg/app"
)

type SendSmsCodeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendSmsCodeService(Context context.Context, RequestContext *app.RequestContext) *SendSmsCodeService {
	return &SendSmsCodeService{RequestContext: RequestContext, Context: Context}
}

func (h *SendSmsCodeService) Run(req *auth2.SendSmsCodeReq) (resp *common.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
