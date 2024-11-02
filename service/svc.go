package service

import (
	"sync"

	"github.com/pkg/errors"

	"github.com/mutezebra/forum/biz/model/forum"
	"github.com/mutezebra/forum/pkg/errno"
	"github.com/mutezebra/forum/repository"
)

type Service struct {
	db *repository.ForumDB
}

var (
	once sync.Once
	obj  *Service
)

func GetService() *Service {
	once.Do(func() {
		obj = &Service{db: repository.NewForumDB()}
	})
	return obj
}

func (svc *Service) VerifyReq(request interface{}) error {
	switch req := request.(type) {
	case *forum.CreateThreadReq:
		if err := svc.verifyTitle(req.GetTitle()); err != nil {
			return err
		}
		if err := svc.verifyContent(req.GetContent()); err != nil {
			return err
		}
	default:
		return errors.New("Unknown req type")
	}

	return nil
}

func (svc *Service) verifyTitle(v string) error {
	if len(v) == 0 || len(v) >= 15 {
		return errno.New(errno.ThreadTitleLengthInvalid, "thread的title长度应该在15个字以内")
	}
	return nil
}

func (svc *Service) verifyContent(v string) error {
	if len(v) == 0 {
		return errno.New(errno.ThreadContentTooShort, "thread的content太短了")
	}
	return nil
}
