package utils

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// SendErrResponse  pack error response
func SendErrResponse(ctx context.Context, c *app.RequestContext, code int, err error) {
	// todo edit custom code
	c.JSON(code, Response{
		Code:    1,
		Message: err.Error(),
	})
}

// SendSuccessResponse  pack success response
func SendSuccessResponse(ctx context.Context, c *app.RequestContext, code int, data interface{}) {
	// todo edit custom code
	c.JSON(code, Response{
		Code:    0,
		Message: "success",
		Data:    data,
	})
}
