package datastore

import (
	"context"
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	"github.com/PaluMacil/barkdognet/datastore/aggregates"
	"github.com/PaluMacil/barkdognet/datastore/types"
	"github.com/go-jet/jet/v2/postgres"
)

// UserProvider is an interface for operations on UserIdentifier.
type UserProvider interface {
	GetUser(ctx context.Context, iden types.UserIdentifier) (*model.SysUser, error)
	All(ctx context.Context, orderBy ...postgres.OrderByClause) ([]model.SysUser, error)
	GetUsersForRole(ctx context.Context, roleID int32) ([]model.SysUser, error)
	SetPassword(ctx context.Context, iden types.UserIdentifier, password string) error
	CheckPassword(ctx context.Context, iden types.UserIdentifier, password string) error
	Update(ctx context.Context, user *model.SysUser) error
	Create(ctx context.Context, user *model.SysUser) error
	Lock(ctx context.Context, iden types.UserIdentifier) error
	UnLock(ctx context.Context, iden types.UserIdentifier) error
	Delete(ctx context.Context, iden types.UserIdentifier) error
}

// RoleProvider is an interface for operations on Role.
type RoleProvider interface {
	Get(ctx context.Context, id int32) (*model.SysRole, error)
	Some(ctx context.Context, iden types.UserIdentifier) ([]model.SysRole, error)
	All(ctx context.Context) ([]model.SysRole, error)
	Update(ctx context.Context, role *model.SysRole) error
	Create(ctx context.Context, role *model.SysRole) error
	AddUser(ctx context.Context, roleID int32, userIden types.UserIdentifier) error
	RemoveUser(ctx context.Context, roleID int32, userIden types.UserIdentifier) error
	Delete(ctx context.Context, id int32) error
}

// SessionProvider is an interface for operations on SessionIdentifier.
type SessionProvider interface {
	GetContext(ctx context.Context, iden types.SessionIdentifier) (aggregates.SessionDetails, error)
	Get(ctx context.Context, iden types.SessionIdentifier) (*model.SysSession, error)
	Create(ctx context.Context, session *model.SysSession) error
	Delete(ctx context.Context, iden types.SessionIdentifier) error
}

// BlogCategoryProvider is an interface for operations on BlogCategory.
type BlogCategoryProvider interface {
	Get(ctx context.Context, id int32) (*model.BlogCategory, error)
	All(ctx context.Context) ([]model.BlogCategory, error)
	Create(ctx context.Context, blogCategory model.BlogCategory) (*model.BlogCategory, error)
	Update(ctx context.Context, blogCategory model.BlogCategory) (*model.BlogCategory, error)
	Delete(ctx context.Context, id model.BlogCategory) error
}

// BlogPostProvider is an interface for operations on BlogPost.
type BlogPostProvider interface {
	Get(ctx context.Context, id int32) (*model.BlogPost, error)
	Some(ctx context.Context, categoryID *int32, page, pageSize int64) ([]model.BlogPost, error)
	GetTags(ctx context.Context, postID int32) ([]model.BlogTag, error)
	All(ctx context.Context, categoryID *int32) ([]model.BlogPost, error)
	Create(ctx context.Context, blogPost *model.BlogPost) error
	Update(ctx context.Context, blogPost *model.BlogPost) error
	Delete(ctx context.Context, id int32) error
}

// BlogCommentProvider is an interface for operations on BlogComment.
type BlogCommentProvider interface {
	Get(ctx context.Context, id int32) (*model.BlogComment, error)
	SomeForBlogPost(ctx context.Context, blogPostID int32, page, pageSize int64) ([]model.BlogComment, error)
	AllForBlogPost(ctx context.Context, blogPostID int32) ([]model.BlogComment, error)
	SomeForUser(ctx context.Context, userIden types.UserIdentifier, page, pageSize int64) ([]model.BlogComment, error)
	Create(ctx context.Context, blogComment *model.BlogComment) error
	Update(ctx context.Context, blogComment *model.BlogComment) error
	Delete(ctx context.Context, id int32) error
}

// BlogCommentLikeProvider is an interface for operations on BlogCommentLike.
type BlogCommentLikeProvider interface {
	Get(ctx context.Context, userID, commentID int32) (*model.BlogCommentLike, error)
	AllForComment(ctx context.Context, commentID int32) ([]model.BlogCommentLike, error)
	CountForComment(ctx context.Context, commentID int32) (int, error)
	Create(ctx context.Context, like *model.BlogCommentLike) error
	Delete(ctx context.Context, userID, commentID int32) error
}

// BlogPostLikeProvider is an interface for operations on BlogPostLike.
type BlogPostLikeProvider interface {
	Get(ctx context.Context, userID, postID int32) (*model.BlogPostLike, error)
	AllForPost(ctx context.Context, postID int32) ([]model.BlogPostLike, error)
	CountForPost(ctx context.Context, postID int32) (int, error)
	Create(ctx context.Context, like *model.BlogPostLike) error
	Delete(ctx context.Context, userID, postID int32) error
}

// BlogTagProvider is an interface for operations on BlogTag.
type BlogTagProvider interface {
	Get(ctx context.Context, id int32) (*model.BlogTag, error)
	AllForPost(ctx context.Context, postID int32) ([]model.BlogTag, error)
	All(ctx context.Context) ([]model.BlogTag, error)
	Create(ctx context.Context, tag *model.BlogTag) error
	Update(ctx context.Context, tag *model.BlogTag) error
	Delete(ctx context.Context, id int32) error
}
