package pgstore

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"
	"github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"
	. "github.com/PaluMacil/barkdognet/.gen/barkdog/public/table"
	"github.com/PaluMacil/barkdognet/datastore"
	"github.com/PaluMacil/barkdognet/datastore/identifier"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
	"log/slog"
)

var _ datastore.RoleProvider = RoleStore{}

type RoleStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

func NewRoleStore(db *stdsql.DB, log *slog.Logger) *RoleStore {
	return &RoleStore{
		db:  db,
		log: log.With(slog.String("DataStore", "RoleStore")),
	}
}

func (r RoleStore) Get(ctx context.Context, id int32) (*model.SysRole, error) {
	var role *model.SysRole
	stmt := SELECT(SysRole.AllColumns).
		FROM(SysRole).
		WHERE(SysRole.ID.EQ(Int32(id)))
	if r.log.Enabled(ctx, slog.LevelDebug) {
		r.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()), slog.Int("id", int(id)))
	}
	if r.log.Enabled(ctx, slog.LevelDebug) {
		r.log.DebugContext(ctx, "Get", slog.Int("id", int(id)), slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, r.db, role)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			r.log.InfoContext(ctx, "Get, not found", slog.Int("id", int(id)))
			return nil, nil
		}
		r.log.ErrorContext(ctx, "Get", slog.Int("id", int(id)), slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting role %d: %w", id, err)
	}
	return role, nil
}

func (r RoleStore) Some(ctx context.Context, iden identifier.User) ([]model.SysRole, error) {
	var roles []model.SysRole
	stmt := SELECT(SysRole.AllColumns).
		FROM(SysRole.
			INNER_JOIN(M2mUserRole, SysRole.ID.EQ(M2mUserRole.SysRoleID)).
			INNER_JOIN(SysUser, SysUser.ID.EQ(M2mUserRole.SysUserID)),
		).
		ORDER_BY(SysRole.DisplayName.ASC())
	if iden.ID != nil {
		stmt = stmt.WHERE(SysUser.ID.EQ(Int32(*iden.ID)))
	} else if iden.Email != nil {
		stmt = stmt.WHERE(SysUser.DisplayName.EQ(String(*iden.Email)))
	} else {
		return nil, identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}
	if r.log.Enabled(ctx, slog.LevelDebug) {
		r.log.DebugContext(ctx, "Some", slog.String("sql", stmt.DebugSql()), iden.Slog())
	}
	err := stmt.QueryContext(ctx, r.db, &roles)
	if err != nil {
		r.log.ErrorContext(ctx, "Some", iden.Slog(), slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting some roles for %s: %w", iden.Slog().String(), err)
	}
	return roles, nil
}

func (r RoleStore) All(ctx context.Context) ([]model.SysRole, error) {
	var roles []model.SysRole
	stmt := SELECT(SysRole.AllColumns).
		FROM(SysRole).
		ORDER_BY(SysRole.DisplayName.ASC())
	if r.log.Enabled(ctx, slog.LevelDebug) {
		r.log.DebugContext(ctx, "All", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, r.db, &roles)
	if err != nil {
		r.log.ErrorContext(ctx, "All", slog.String("err", err.Error()))
		return nil, fmt.Errorf("getting all roles: %w", err)
	}
	return roles, nil
}

func (r RoleStore) Update(ctx context.Context, role *model.SysRole) error {
	stmt := SysRole.UPDATE(SysRole.MutableColumns.Except(SysRole.CreatedAt)).
		MODEL(role).
		WHERE(SysRole.ID.EQ(Int32(role.ID)))
	if r.log.Enabled(ctx, slog.LevelDebug) {
		r.log.DebugContext(ctx, "Update", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, r.db)
	if err != nil {
		r.log.ErrorContext(ctx, "Update", slog.String("err", err.Error()))
		return fmt.Errorf("updating role: %w", err)
	}
	return nil
}

func (r RoleStore) Create(ctx context.Context, role *model.SysRole) error {
	stmt := SysRole.INSERT(SysRole.MutableColumns.Except(SysRole.CreatedAt)).
		MODEL(role).
		RETURNING(SysRole.AllColumns)
	if r.log.Enabled(ctx, slog.LevelDebug) {
		r.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, r.db, role)
	if err != nil {
		r.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("creating role: %w", err)
	}
	return nil
}

func (r RoleStore) AddUser(ctx context.Context, roleID int32, userIden identifier.User) error {
	if userIden.ID == nil {
		if userIden.Email == nil {
			id := fmt.Sprintf("roleID %d, user %s", roleID, userIden.Slog().String())
			return identifier.ErrInsufficient{IdentString: id}
		}
		stmt2 := SELECT(SysUser.ID).
			FROM(SysUser).
			WHERE(SysUser.Email.EQ(String(*userIden.Email)))
		err := stmt2.QueryContext(ctx, r.db, userIden.ID)
		if err != nil {
			return fmt.Errorf("while adding user to role, getting id from email: %w", err)
		}
	}
	stmt := M2mUserRole.INSERT(SysRole.ID, SysUser.ID)
	stmt = stmt.VALUES(roleID, userIden.ID)
	_, err := stmt.ExecContext(ctx, r.db)
	if err != nil {
		return fmt.Errorf("adding user to role: %w", err)
	}
	return nil
}

func (r RoleStore) RemoveUser(ctx context.Context, roleID int32, userIden identifier.User) error {

	if userIden.ID == nil {
		if userIden.Email == nil {
			id := fmt.Sprintf("roleID %d, user %s", roleID, userIden.Slog().String())
			return identifier.ErrInsufficient{IdentString: id}
		}
		stmt2 := SELECT(SysUser.ID).
			FROM(SysUser).
			WHERE(SysUser.Email.EQ(String(*userIden.Email)))
		err := stmt2.QueryContext(ctx, r.db, userIden.ID)
		if err != nil {
			return fmt.Errorf("while removing user from role, getting id from email: %w", err)
		}
	}
	stmt := M2mUserRole.DELETE().
		WHERE(AND(
			M2mUserRole.SysRoleID.EQ(Int32(roleID)),
			M2mUserRole.SysUserID.EQ(Int32(*userIden.ID)),
		))
	_, err := stmt.ExecContext(ctx, r.db)
	if err != nil {
		return fmt.Errorf("removing user from role: %w", err)
	}
	return nil
}

func (r RoleStore) Delete(ctx context.Context, id int32) error {
	stmt := SysRole.DELETE().WHERE(SysRole.ID.EQ(Int32(id)))
	if r.log.Enabled(ctx, slog.LevelDebug) {
		r.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, r.db)
	if err != nil {
		r.log.ErrorContext(ctx, "Delete", slog.String("err", err.Error()))
		return fmt.Errorf("deleting role %d: %w", id, err)
	}
	return nil
}
