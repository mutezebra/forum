namespace go forum

include "base.thrift"

struct ForumThread {
    1: optional i64 ID (api.body="id")
    2: optional string Title (api.body="title")
    3: optional i64 CreatorID (api.body="creator_id")
    4: optional string Content (api.body="content")
    5: optional i64 CreatedAt (api.body="created_at")
    6: optional i64 UpdatedAt (api.body="updated_at")
}

struct ForumPost {
    1: optional i64 ID (api.body="id")
    2: optional i64 ThreadID (api.body="thread_id")
    3: optional i64 CreatorID (api.body="creator_id")
    4: optional string Content (api.body="content")
    5: optional i64 CreatedAt (api.body="created_at")
    6: optional i64 UpdatedAt (api.body="updated_at")
}

struct ForumComment {
    1: optional i64 ID (api.body="id")
    2: optional i64 PostID (api.body="post_id")
    3: optional i64 CreatorID (api.body="creator_id")
    4: optional string Content (api.body="content")
    5: optional i64 CreatedAt (api.body="created_at")
    6: optional i64 UpdatedAt (api.body="updated_at")
}

struct GetThreadsReq {
    1: optional i64 Pages (api.form="pages")
    2: optional i64 Size (api.form="size")
}

struct GetThreadsResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 Total (api.body="total")
    3: optional list<ForumThread> Threads (api.body="threads")
}

struct GetPostsReq {
    1: optional i64 ThreadID (api.query="thread_id", api.path="thread_id")
    2: optional i64 Pages (api.form="pages")
    3: optional i64 Size (api.form="size")
}

struct GetPostsResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 Total (api.body="total")
    3: optional list<ForumPost> Posts (api.body="posts")
}

struct GetCommentsReq {
    1: optional i64 PostID (api.query="post_id", api.path="post_id")
    2: optional i64 Pages (api.form="pages")
    3: optional i64 Size (api.form="size")
}

struct GetCommentsResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 Total (api.body="total")
    3: optional list<ForumComment> Comments (api.body="comments")
}

struct CreateThreadReq {
    1: optional i64 UserID
    2: optional string Title (api.form="title")
    3: optional string Content (api.form="content")
}

struct CreateThreadResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 ThreadID (api.body="thread_id")
}

struct CreatePostReq {
    1: optional i64 UserID
    2: optional i64 ThreadID (api.form="thread_id")
    3: optional string Content (api.form="content")
}

struct CreatePostResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 PostID (api.body="post_id")
}

struct CreateCommentReq {
    1: optional i64 UserID
    2: optional i64 PostID (api.form="post_id")
    3: optional string Content (api.form="content")
}

struct CreateCommentResp {
    1: optional base.Base Base (api.body="base")
    2: optional i64 CommentID (api.body="comment_id")
}

service ForumService {
    GetThreadsResp GetThreads(1: GetThreadsReq req) (api.post="/api/forum/get-threads")
    GetPostsResp GetPosts(1: GetPostsReq req) (api.post="/api/forum/get-posts")
    GetCommentsResp GetComments(1: GetCommentsReq req) (api.post="/api/forum/get-comments")
    CreateThreadResp CreateThread(1: CreateThreadReq req) (api.post="/api/auth/forum/create-thread")
    CreatePostResp CreatePost(1: CreatePostReq req) (api.post="/api/auth/forum/create-post")
    CreateCommentResp CreateComment(1: CreateCommentReq req) (api.post="/api/auth/forum/create-comment")
}

