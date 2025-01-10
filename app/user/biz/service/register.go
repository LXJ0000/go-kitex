package service

import (
	"context"

	"github.com/LXJ0000/go-kitex/app/user/biz/dal/mysql"
	"github.com/LXJ0000/go-kitex/app/user/model"
	user "github.com/LXJ0000/go-kitex/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	defer func() {
		klog.Info("RegisterService", "req", req, "resp", resp, "err", err)
	}()
	// Finish your business logic.
	item := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := model.CreateUser(mysql.DB, item); err != nil {
		return nil, err
	}
	resp = &user.RegisterResp{
		UserId: int32(item.ID),
	}
	return
}
