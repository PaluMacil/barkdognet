package datastore

import "github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"

type UserIdentifier struct {
	ID    *int32
	Email *string
}

// UserProvider provides User retrieval and manipulation
type UserProvider interface {
	GetUser(iden UserIdentifier) (model.SysUser, error)
	GetUsersForRole(roleID int32) ([]model.SysUser, error)
	CreateUser(user model.SysUser) (model.SysUser, error)
}

// RoleProvider provides Role retrieval and manipulation
type RoleProvider interface {
	GetRoleByID(id int32) (model.SysRole, error)
	GetRolesForUser(iden UserIdentifier) (model.SysRole, error)
	AddUser(userIden UserIdentifier, roleID int32) error
	RemoveUser(roleIdentity, userId int32) error
}
