package pgstore

import (
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	"github.com/PaluMacil/barkdognet/datastore"
	"github.com/PaluMacil/barkdognet/datastore/identifier"
)

var _ datastore.BlogCommentProvider = BlogCommentStore{}

type BlogCommentStore struct{}

func (b BlogCommentStore) Get(id int32) (model.BlogComment, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCommentStore) SomeForBlogPost(blogPostID int32, page, pageSize, number int) ([]model.BlogComment, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCommentStore) AllForBlogPost(blogPostID *int32) ([]model.BlogComment, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCommentStore) SomeForUser(userIden identifier.User, page, pageSize, number int) ([]model.BlogComment, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCommentStore) Delete(id int32) error {
	//TODO implement me
	panic("implement me")
}
