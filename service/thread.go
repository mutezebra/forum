package service

import (
	"context"
	"github.com/mutezebra/forum/biz/model/forum"
	"github.com/mutezebra/forum/pkg/pack"
	"github.com/mutezebra/forum/pkg/utils"
	"time"
)

func (svc *Service) CreateThread(ctx context.Context, req *forum.CreateThreadReq) (resp *forum.CreateThreadResp, err error) {
	defer func() {
		pack.LogError(err)
	}()

	if err = svc.VerifyReq(req); err != nil {
		return nil, err
	}

	thread := buildThread(req.GetUserID(), req.GetTitle(), req.GetContent())
	if err = svc.db.CreateThread(ctx, thread); err != nil {
		return nil, err
	}

	resp = new(forum.CreateThreadResp)
	resp.Base = pack.Base
	return resp, nil
}

func buildThread(uid int64, title, content string) *forum.ForumThread {
	return &forum.ForumThread{
		Title:     utils.Ptr(title),
		CreatorID: utils.Ptr(uid),
		Content:   utils.Ptr(content),
		CreatedAt: utils.Ptr(time.Now().Unix()),
		UpdatedAt: utils.Ptr(time.Now().Unix()),
	}
}
