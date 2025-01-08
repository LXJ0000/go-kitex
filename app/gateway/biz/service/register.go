package service

import (
	"context"

	auth "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/gateway/auth"
	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth.RegisterReq) (resp *auth.RegisterResp, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
