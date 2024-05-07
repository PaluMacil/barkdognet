package pgstore

import (
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	"github.com/PaluMacil/barkdognet/datastore"
)

var _ datastore.BlogPostProvider = BlogPostStore{}

type BlogPostStore struct{}

func (b BlogPostStore) Get(id int32) (model.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostStore) Some(categoryID *int32, page, pageSize, number int) ([]model.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostStore) All(categoryID *int32) ([]model.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostStore) Create(blogPost model.BlogPost) (model.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostStore) Update(blogPost model.BlogPost) (model.BlogPost, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogPostStore) Delete(id int32) error {
	//TODO implement me
	panic("implement me")
}
