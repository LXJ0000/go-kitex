package service

import (
	"context"

	"github.com/LXJ0000/go-kitex/app/user/biz/dal/mysql"
	"github.com/LXJ0000/go-kitex/app/user/model"
	user "github.com/LXJ0000/go-kitex/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	item := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	if err := model.CreateUser(mysql.DB, item); err != nil {
		return nil, err
	}
	return &user.LoginResp{
		UserId: int32(item.ID),
	}, nil
}
