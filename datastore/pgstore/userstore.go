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
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"strconv"
)

// UserStore is a data access layer for user data in Postgres.
var _ datastore.UserProvider = UserStore{}

type UserStore struct {
	db  *stdsql.DB
	log *slog.Logger
}

// NewUserStore creates a new UserStore with the given database and logger.
func NewUserStore(db *stdsql.DB, log *slog.Logger) *UserStore {
	return &UserStore{
		db:  db,
		log: log.With(slog.String("DataStore", "UserStore")),
	}
}

// GetUser retrieves a user from the database using their identifier.
func (u UserStore) GetUser(ctx context.Context, iden identifier.User) (*model.SysUser, error) {
	var user *model.SysUser
	stmt := SELECT(SysUser.AllColumns).FROM(SysUser)
	if iden.ID != nil {
		stmt = stmt.WHERE(SysUser.ID.EQ(Int32(*iden.ID)))
	} else if iden.Email != nil {
		stmt = stmt.WHERE(SysUser.Email.EQ(String(*iden.Email)))
	} else {
		return nil, identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "GetUser", slog.String("sql", stmt.DebugSql()), iden.Slog())
	}
	err := stmt.QueryContext(ctx, u.db, user)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			u.log.InfoContext(ctx, "GetUser, not found", iden.Slog())
			return nil, nil
		}
		u.log.ErrorContext(ctx, "GetUser", iden.Slog(), slog.String("err", err.Error()))
		return nil, fmt.Errorf("error getting user %s: %w", iden.Slog(), err)
	}
	return user, nil
}

// GetUsersForRole retrieves a list of users associated with a specific role ID.
func (u UserStore) GetUsersForRole(ctx context.Context, roleID int32) ([]model.SysUser, error) {
	var users []model.SysUser
	stmt := SELECT(SysUser.AllColumns).
		FROM(SysUser.
			INNER_JOIN(M2mUserRole, SysUser.ID.EQ(M2mUserRole.SysUserID)).
			INNER_JOIN(SysRole, SysRole.ID.EQ(M2mUserRole.SysRoleID)),
		).
		WHERE(SysRole.ID.EQ(Int32(roleID))).
		ORDER_BY(SysRole.DisplayName)
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "GetUsersForRole", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, u.db, &users)
	if err != nil {
		u.log.ErrorContext(ctx, "GetUsersForRole", slog.Int("roleID", int(roleID)), slog.String("err", err.Error()))
		return users, fmt.Errorf("error getting users for roleID %s: %w", strconv.Itoa(int(roleID)), err)
	}
	return users, nil
}

// SetPassword sets the password of a user identified by their identifier.
func (u UserStore) SetPassword(ctx context.Context, iden identifier.User, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		u.log.ErrorContext(ctx, "SetPassword, hash generation", slog.String("err", err.Error()))
		return fmt.Errorf("error generating hash for setting password: %w", err)
	}
	stmt := SysUser.UPDATE(SysUser.PasswordHash).SET(hash)
	if iden.ID != nil {
		stmt = stmt.WHERE(SysUser.ID.EQ(Int32(*iden.ID)))
	} else if iden.Email != nil {
		stmt = stmt.WHERE(SysUser.Email.EQ(String(*iden.Email)))
	} else {
		return identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "SetPassword", iden.Slog(), slog.String("sql", stmt.DebugSql()))
	}
	_, err = stmt.ExecContext(ctx, u.db)
	if err != nil {
		u.log.ErrorContext(ctx, "SetPassword, executing update", iden.Slog(), slog.String("err", err.Error()))
		return fmt.Errorf("error executing update command for setting password for %s: %w", iden.Slog(), err)
	}
	return nil
}

// CheckPassword checks whether the provided password matches the stored one for a particular user.
func (u UserStore) CheckPassword(ctx context.Context, iden identifier.User, password string) error {
	stmt := SELECT(SysUser.PasswordHash).FROM(SysUser)
	if iden.ID != nil {
		stmt = stmt.WHERE(SysUser.ID.EQ(Int32(*iden.ID)))
	} else if iden.Email != nil {
		stmt = stmt.WHERE(SysUser.Email.EQ(String(*iden.Email)))
	} else {
		return identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "CheckPassword", slog.String("sql", stmt.DebugSql()), iden.Slog())
	}
	var hash []byte
	err := stmt.QueryContext(ctx, u.db, hash)
	if err != nil {
		u.log.ErrorContext(ctx, "CheckPassword, error querying hash", slog.String("err", err.Error()), iden.Slog())
		return fmt.Errorf("error querying for hash check%d: %w", err)
	}
	if err := bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		u.log.InfoContext(ctx, "CheckPassword, error comparing hash: %s", err.Error())
		return fmt.Errorf("error checking password for %s: %w", iden.Slog(), err)
	}
	return nil
}

// Update updates a User in the database
func (u UserStore) Update(ctx context.Context, user *model.SysUser) error {
	stmt := SysUser.UPDATE(SysUser.MutableColumns.Except(SysUser.CreatedAt)).
		MODEL(user).
		WHERE(SysUser.ID.EQ(Int32(user.ID)))
	if u.log.Enabled(ctx, slog.LevelDebug) {
		// TODO: log identity after adding the method to get one from a model
		u.log.DebugContext(ctx, "UpdateUser", slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, u.db)
	if err != nil {
		u.log.ErrorContext(ctx, "Update", slog.String("err", err.Error()))
		return fmt.Errorf("updating user: %w", err)
	}
	return nil
}

// Create creates a new user in the system.
func (u UserStore) Create(ctx context.Context, user *model.SysUser) error {
	stmt := SysUser.INSERT(SysUser.MutableColumns.Except(SysUser.CreatedAt, SysUser.PasswordHash)).
		MODEL(user).
		RETURNING(SysUser.AllColumns)
	if u.log.Enabled(ctx, slog.LevelDebug) {
		// TODO: log identity after adding the method to get one from a model
		u.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, u.db, user)
	if err != nil {
		u.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("creating user: %w", err)
	}
	return nil
}

// Lock locks a user account. This can be useful in case the user made several invalid login attempts.
func (u UserStore) Lock(ctx context.Context, iden identifier.User) error {
	stmt := SysUser.UPDATE(SysUser.Locked).SET(true)
	if iden.ID != nil {
		stmt = stmt.WHERE(SysUser.ID.EQ(Int32(*iden.ID)))
	} else if iden.Email != nil {
		stmt = stmt.WHERE(SysUser.Email.EQ(String(*iden.Email)))
	} else {
		return identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "Lock", iden.Slog(), slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, u.db)
	if err != nil {
		u.log.ErrorContext(ctx, "Lock", iden.Slog(), slog.String("err", err.Error()))
		return fmt.Errorf("lock, executing update for %s: %w", iden.Slog(), err)
	}
	return nil
}

// UnLock unlocks a user account that has been locked previously.
func (u UserStore) UnLock(ctx context.Context, iden identifier.User) error {
	stmt := SysUser.UPDATE(SysUser.Locked).SET(false)
	if iden.ID != nil {
		stmt = stmt.WHERE(SysUser.ID.EQ(Int32(*iden.ID)))
	} else if iden.Email != nil {
		stmt = stmt.WHERE(SysUser.Email.EQ(String(*iden.Email)))
	} else {
		return identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "UnLock", iden.Slog(), slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, u.db)
	if err != nil {
		u.log.ErrorContext(ctx, "UnLock", iden.Slog(), slog.String("err", err.Error()))
		return fmt.Errorf("unlock, executing update for %s: %w", iden.Slog(), err)
	}
	return nil

}

// All returns all the users present in the system.
// Optionally, the returned users list can be ordered according to the provided order by clause.
func (u UserStore) All(ctx context.Context, orderBy ...OrderByClause) ([]model.SysUser, error) {
	var users []model.SysUser
	if len(orderBy) == 0 {
		orderBy = []OrderByClause{SysUser.DisplayName.ASC()}
	}
	stmt := SELECT(SysUser.AllColumns).FROM(SysUser).ORDER_BY(orderBy...)
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "All", slog.String("sql", stmt.DebugSql()))
	}
	err := stmt.QueryContext(ctx, u.db, &users)
	if err != nil {
		u.log.ErrorContext(ctx, "All", slog.String("err", err.Error()))
		return users, fmt.Errorf("getting all users, executing sql: %w", err)
	}
	return users, nil
}

// Delete deletes a user account from the system.
func (u UserStore) Delete(ctx context.Context, iden identifier.User) error {
	stmt := SysUser.DELETE()
	if iden.ID != nil {
		stmt = stmt.WHERE(SysUser.ID.EQ(Int32(*iden.ID)))
	} else if iden.Email != nil {
		stmt = stmt.WHERE(SysUser.Email.EQ(String(*iden.Email)))
	} else {
		return identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}
	if u.log.Enabled(ctx, slog.LevelDebug) {
		u.log.DebugContext(ctx, "Delete", iden.Slog(), slog.String("sql", stmt.DebugSql()))
	}
	_, err := stmt.ExecContext(ctx, u.db)
	if err != nil {
		u.log.ErrorContext(ctx, "Delete", iden.Slog(), slog.String("err", err.Error()))
		return fmt.Errorf("deleting user %s: %w", iden.Slog(), err)
	}
	return nil
}
