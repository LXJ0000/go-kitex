package service

import (
	"context"

	auth2 "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/gateway/auth2"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginSmsCodeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginSmsCodeService(Context context.Context, RequestContext *app.RequestContext) *LoginSmsCodeService {
	return &LoginSmsCodeService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginSmsCodeService) Run(req *auth2.LoginSmsCodeReq) (resp *auth2.LoginSmsCodeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
