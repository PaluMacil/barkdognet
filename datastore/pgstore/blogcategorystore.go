package pgstore

import (
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	"github.com/PaluMacil/barkdognet/datastore"
)

var _ datastore.BlogCategoryProvider = BlogCategoryStore{}

type BlogCategoryStore struct{}

func (b BlogCategoryStore) Get(id int32) (model.BlogCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCategoryStore) All() ([]model.BlogCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCategoryStore) Create(blogCategory model.BlogCategory) (model.BlogCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCategoryStore) Update(blogCategory model.BlogCategory) (model.BlogCategory, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlogCategoryStore) Delete(id model.BlogCategory) error {
	//TODO implement me
	panic("implement me")
}
