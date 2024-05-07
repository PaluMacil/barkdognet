package pgstore

import (
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	"github.com/PaluMacil/barkdognet/datastore"
	"github.com/PaluMacil/barkdognet/datastore/identifier"
)

var _ datastore.SessionProvider = SessionStore{}

type SessionStore struct{}

func (s SessionStore) GetContext(iden identifier.Session) (*identifier.SessionContext, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionStore) Get(iden identifier.Session) (model.SysSession, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionStore) All() ([]model.SysSession, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionStore) Create(session model.SysSession) (model.SysSession, error) {
	//TODO implement me
	panic("implement me")
}

func (s SessionStore) Delete(iden identifier.Session) error {
	//TODO implement me
	panic("implement me")
}
