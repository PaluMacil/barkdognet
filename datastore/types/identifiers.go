package types

import (
	"fmt"
	"log/slog"
)

type ErrInsufficient struct {
	IdentString string
}

func (ei ErrInsufficient) Error() string {
	return fmt.Sprintf("insufficient identity specified: %s", ei.IdentString)
}

type UserIdentifier struct {
	ID    *int32
	Email *string
}

func (i UserIdentifier) Slog() slog.Attr {
	if i.ID != nil {
		return slog.Int("Identifier", int(*i.ID))
	}
	if i.Email != nil {
		return slog.String("Identifier", "Email"+*i.Email)
	}
	return slog.String("Identifier", "None Specified")
}

type SessionIdentifier struct {
	ID    *int32
	Token *string
}

func (i SessionIdentifier) Slog() slog.Attr {
	if i.ID != nil {
		return slog.Int("ID", int(*i.ID))
	}
	if i.Token != nil {
		if len(*i.Token) > 4 {
			return slog.String("Identifier", "Token"+*i.Token+"****")
		}
		return slog.String("Identifier", "Token"+"????")
	}
	return slog.String("Identifier", "None Specified")
}
