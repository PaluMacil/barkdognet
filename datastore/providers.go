package datastore

import (
	"context"
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	"github.com/PaluMacil/barkdognet/datastore/identifier"
	"github.com/go-jet/jet/v2/postgres"
)

// UserProvider provides User retrieval and manipulation
type UserProvider interface {
	GetUser(ctx context.Context, iden identifier.User) (*model.SysUser, error)
	All(ctx context.Context, orderBy ...postgres.OrderByClause) ([]model.SysUser, error)
	GetUsersForRole(ctx context.Context, roleID int32) ([]model.SysUser, error)
	SetPassword(ctx context.Context, iden identifier.User, password string) error
	CheckPassword(ctx context.Context, iden identifier.User, password string) error
	Update(ctx context.Context, user *model.SysUser) error
	Create(ctx context.Context, user *model.SysUser) error
	Lock(ctx context.Context, iden identifier.User) error
	UnLock(ctx context.Context, iden identifier.User) error
	Delete(ctx context.Context, iden identifier.User) error
}

// RoleProvider provides Role retrieval and manipulation
type RoleProvider interface {
	Get(ctx context.Context, id int32) (*model.SysRole, error)
	Some(ctx context.Context, iden identifier.User) ([]model.SysRole, error)
	All(ctx context.Context) ([]model.SysRole, error)
	Update(ctx context.Context, role *model.SysRole) error
	Create(ctx context.Context, role *model.SysRole) error
	AddUser(ctx context.Context, roleID int32, userIden identifier.User) error
	RemoveUser(ctx context.Context, roleID int32, userIden identifier.User) error
	Delete(ctx context.Context, id int32) error
}

type SessionProvider interface {
	GetContext(ctx context.Context, iden identifier.Session) (identifier.SessionContext, error)
	Get(ctx context.Context, iden identifier.Session) (*model.SysSession, error)
	Create(ctx context.Context, session *model.SysSession) error
	Delete(ctx context.Context, iden identifier.Session) error
}

type BlogCategoryProvider interface {
	Get(ctx context.Context, id int32) (*model.BlogCategory, error)
	All(ctx context.Context) ([]model.BlogCategory, error)
	Create(ctx context.Context, blogCategory model.BlogCategory) (*model.BlogCategory, error)
	Update(ctx context.Context, blogCategory model.BlogCategory) (*model.BlogCategory, error)
	Delete(ctx context.Context, id model.BlogCategory) error
}

type BlogPostProvider interface {
	Get(ctx context.Context, id int32) (*model.BlogPost, error)
	Some(ctx context.Context, categoryID *int32, page, pageSize int64) ([]model.BlogPost, error)
	All(ctx context.Context, categoryID *int32) ([]model.BlogPost, error)
	Create(ctx context.Context, blogPost *model.BlogPost) error
	Update(ctx context.Context, blogPost *model.BlogPost) error
	Delete(ctx context.Context, id int32) error
}

type BlogCommentProvider interface {
	Get(ctx context.Context, id int32) (*model.BlogComment, error)
	SomeForBlogPost(ctx context.Context, blogPostID int32, page, pageSize int64) ([]model.BlogComment, error)
	AllForBlogPost(ctx context.Context, blogPostID int32) ([]model.BlogComment, error)
	SomeForUser(ctx context.Context, userIden identifier.User, page, pageSize int64) ([]model.BlogComment, error)
	Create(ctx context.Context, blogComment *model.BlogComment) error
	Update(ctx context.Context, blogComment *model.BlogComment) error
	Delete(ctx context.Context, id int32) error
}
