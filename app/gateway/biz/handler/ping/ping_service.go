package ping

import (
	"context"

	"github.com/LXJ0000/go-kitex/app/gateway/biz/service"
	"github.com/LXJ0000/go-kitex/app/gateway/biz/utils"
	common "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/common"
	ping "github.com/LXJ0000/go-kitex/app/gateway/hertz_gen/gateway/ping"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// Ping .
// @router /ping [GET]
func Ping(ctx context.Context, c *app.RequestContext) {
	var err error
	var req common.Empty
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &ping.PingResponse{}
	resp, err = service.NewPingService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
