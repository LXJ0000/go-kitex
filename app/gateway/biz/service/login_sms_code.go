package service

import (
	"context"

	auth "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type LoginSmsCodeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginSmsCodeService(Context context.Context, RequestContext *app.RequestContext) *LoginSmsCodeService {
	return &LoginSmsCodeService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginSmsCodeService) Run(req *auth.LoginSmsCodeReq) (resp *auth.LoginSmsCodeResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
