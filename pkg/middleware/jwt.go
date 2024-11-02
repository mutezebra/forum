package middleware

import (
	"context"
	"github.com/mutezebra/forum/pkg/constants"
	"github.com/mutezebra/forum/pkg/errno"
	"github.com/mutezebra/forum/pkg/jwt"
	"github.com/mutezebra/forum/pkg/pack"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func JWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := string(c.GetHeader(constants.TokenHeader))
		if token == "" {
			pack.SendErrResp(c, errno.New(errno.LackToken, "lack of token"), consts.StatusUnauthorized)
			c.Abort()
			return
		}

		uid, valid, err := jwt.CheckToken(token)
		if uid == "" {
			pack.SendErrResp(c, errno.New(errno.WrongToken, err.Error()), consts.StatusUnauthorized)
			c.Abort()
			return
		}
		if !valid || err != nil {
			pack.SendErrResp(c, errno.New(errno.TokenExpire, "token is expired, please log in again"), consts.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Next(context.WithValue(ctx, constants.TokenKey, uid))
	}
}
