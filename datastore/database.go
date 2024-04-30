package datastore

type Database struct {
	Users UserProvider
	Roles RoleProvider
}
