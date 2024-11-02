package pack

import (
	"errors"
	"github.com/mutezebra/forum/biz/model/base"
	"github.com/mutezebra/forum/pkg/errno"
	"github.com/mutezebra/forum/pkg/logger"
	"github.com/mutezebra/forum/pkg/utils"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

var Base *base.Base

func init() {
	Base = &base.Base{
		Code: utils.Ptr(int32(200)),
		Msg:  utils.Ptr("ok"),
	}
}

func LogError(err error) {
	if err == nil {
		return
	}

	var e errno.Errno
	if errors.As(err, &e) {
		logger.KDebugf("Code: %d, Message: %s", e.Code(), e.Error())
		return
	}
	logger.KErrorf("Message: %+v", err)
}

func SendErrResp(c *app.RequestContext, err error, codes ...int) {
	var e errno.Errno
	if errors.As(err, &e) {
		code := consts.StatusOK
		if codes != nil {
			code = codes[0]
		}

		c.JSON(code, base.BaseResp{Base: &base.Base{
			Code: utils.Ptr(e.Code()),
			Msg:  utils.Ptr(e.Error()),
		}})
		return
	}

	c.JSON(consts.StatusInternalServerError, base.BaseResp{Base: &base.Base{
		Code: utils.Ptr(int32(consts.StatusInternalServerError)),
		Msg:  utils.Ptr(err.Error()),
	}})
}
