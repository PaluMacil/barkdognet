package pgstore

import (
	"context"
	"database/sql"
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

var _ datastore.SessionProvider = SessionStore{}

type SessionStore struct {
	db  *sql.DB
	log *slog.Logger
}

func NewSessionStore(db *sql.DB, log *slog.Logger) *SessionStore {
	return &SessionStore{
		db:  db,
		log: log.With(slog.String("DataStore", "SessionStore")),
	}
}

func (s SessionStore) GetContext(ctx context.Context, iden identifier.Session) (identifier.SessionContext, error) {
	sessionContext := identifier.SessionContext{}
	stmt := SELECT(SysSession.AllColumns, SysUser.AllColumns, SysRole.AllColumns).
		FROM(SysSession.INNER_JOIN(SysUser, SysSession.SysUserID.EQ(SysUser.ID)).
			INNER_JOIN(M2mUserRole, SysUser.ID.EQ(M2mUserRole.SysUserID)).
			INNER_JOIN(SysRole, M2mUserRole.SysRoleID.EQ(SysRole.ID))).
		WHERE(SysSession.SessionToken.EQ(String(*iden.Token)))

	if s.log.Enabled(ctx, slog.LevelDebug) {
		s.log.DebugContext(ctx, "GetContext", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, s.db, &sessionContext)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			s.log.InfoContext(ctx, "GetContext, not found", iden.Slog())
			return identifier.SessionContext{}, nil
		}
		s.log.ErrorContext(ctx, "GetContext", slog.String("err", err.Error()))
		return identifier.SessionContext{}, fmt.Errorf("error getting session context for token %s: %w", iden.Token, err)
	}

	return sessionContext, nil
}

func (s SessionStore) Get(ctx context.Context, iden identifier.Session) (*model.SysSession, error) {
	var session model.SysSession
	stmt := SELECT(SysSession.AllColumns).FROM(SysSession)
	if iden.ID != nil {
		stmt = stmt.WHERE(SysSession.ID.EQ(Int32(*iden.ID)))
	} else if iden.Token != nil {
		stmt = stmt.WHERE(SysSession.SessionToken.EQ(String(*iden.Token)))
	} else {
		return nil, identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}

	if s.log.Enabled(ctx, slog.LevelDebug) {
		s.log.DebugContext(ctx, "Get", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, s.db, &session)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			s.log.InfoContext(ctx, "Get, not found", iden.Slog())
			return nil, nil
		}
		s.log.ErrorContext(ctx, "Get", slog.String("err", err.Error()))
		return nil, fmt.Errorf("error getting session %d: %w", iden.ID, err)
	}
	return &session, nil
}

func (s SessionStore) Create(ctx context.Context, session *model.SysSession) error {
	stmt := SysSession.INSERT(SysSession.MutableColumns).
		MODEL(session).
		RETURNING(SysSession.AllColumns)

	if s.log.Enabled(ctx, slog.LevelDebug) {
		s.log.DebugContext(ctx, "Create", slog.String("sql", stmt.DebugSql()))
	}

	err := stmt.QueryContext(ctx, s.db, session)
	if err != nil {
		s.log.ErrorContext(ctx, "Create", slog.String("err", err.Error()))
		return fmt.Errorf("error creating session: %w", err)
	}
	return nil
}

func (s SessionStore) Delete(ctx context.Context, iden identifier.Session) error {
	stmt := SysSession.DELETE()
	if iden.ID != nil {
		stmt = stmt.WHERE(SysSession.ID.EQ(Int32(*iden.ID)))
	} else if iden.Token != nil {
		stmt = stmt.WHERE(SysSession.SessionToken.EQ(String(*iden.Token)))
	} else {
		return identifier.ErrInsufficient{IdentString: iden.Slog().String()}
	}

	if s.log.Enabled(ctx, slog.LevelDebug) {
		s.log.DebugContext(ctx, "Delete", slog.String("sql", stmt.DebugSql()))
	}

	_, err := stmt.ExecContext(ctx, s.db)
	if err != nil {
		s.log.ErrorContext(ctx, "Delete", slog.String("err", err.Error()))
		return fmt.Errorf("error deleting session %d: %w", iden.ID, err)
	}
	return nil
}
