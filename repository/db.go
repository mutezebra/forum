package repository

import (
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/mutezebra/forum/biz/model/forum"
	"github.com/mutezebra/forum/pkg/client"
)

type ForumDB struct {
	db *gorm.DB
}

const (
	ThreadTable  = "forum_threads"
	PostsTable   = "forum_posts"
	CommentTable = "forum_comments"
)

func NewForumDB() *ForumDB {
	return &ForumDB{client.GetMysqlDB()}
}

func (db *ForumDB) CreateThread(ctx context.Context, thread *forum.ForumThread) error {
	if err := db.db.Table(ThreadTable).Create(&thread).Error; err != nil {
		return errors.Errorf("failed when create thread,err: %v", err)
	}
	return nil
}
