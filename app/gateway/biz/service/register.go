package service

import (
	"context"

	auth2 "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/gateway/auth2"
	"github.com/LXJ0000/go-kitex/app/gateway/infra/rpc"
	"github.com/LXJ0000/go-kitex/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type RegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterService(Context context.Context, RequestContext *app.RequestContext) *RegisterService {
	return &RegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterService) Run(req *auth2.RegisterReq) (resp *auth2.RegisterResp, err error) {
	defer func() {
		hlog.CtxInfof(h.Context, "req = %+v", req)
		hlog.CtxInfof(h.Context, "resp = %+v", resp)
	}()
	r, err := rpc.UserClient.Register(h.Context, &user.RegisterReq{
		Email:           req.Email,
		Password:        req.Password,
		ConfirmPassword: req.ConfirmPassword,
	})
	if err != nil {
		return nil, err
	}
	resp = &auth2.RegisterResp{
		UserId: r.UserId,
	}
	return 
}
