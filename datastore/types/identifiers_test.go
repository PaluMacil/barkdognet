package types_test

import (
	"github.com/PaluMacil/barkdognet/datastore/types"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestErrInsufficient_Error(t *testing.T) {
	tests := map[string]struct {
		err      types.ErrInsufficient
		expected string
	}{
		"NoIdentity": {
			err:      types.ErrInsufficient{IdentString: ""},
			expected: "insufficient identity specified: ",
		},
		"WithIdentity": {
			err:      types.ErrInsufficient{IdentString: "username"},
			expected: "insufficient identity specified: username",
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			assert.Equal(t, testCase.expected, testCase.err.Error())
		})
	}
}

func TestUserIdentifier_Slog(t *testing.T) {
	tests := map[string]struct {
		id       types.UserIdentifier
		expected slog.Attr
	}{
		"IdentifierWithID": {
			id:       types.UserIdentifier{ID: PInt32(42), Email: PString("test@example.com")},
			expected: slog.Int("Identifier", 42),
		},
		"IdentifierWithEmail": {
			id:       types.UserIdentifier{Email: PString("test@example.com")},
			expected: slog.String("Identifier", "Emailtest@example.com"),
		},
		"IdentifierWithoutIDOrEmail": {
			id:       types.UserIdentifier{},
			expected: slog.String("Identifier", "None Specified"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actualAttr := testCase.id.Slog()
			assert.Equal(t, testCase.expected, actualAttr)
		})
	}
}

func TestSessionIdentifier_Slog(t *testing.T) {
	tests := map[string]struct {
		id       types.SessionIdentifier
		expected slog.Attr
	}{
		"IdentifierWithID": {
			id:       types.SessionIdentifier{ID: PInt32(42), Token: PString("abcdef")},
			expected: slog.Int("ID", 42),
		},
		"IdentifierWithLongToken": {
			id:       types.SessionIdentifier{Token: PString("abcdef")},
			expected: slog.String("Identifier", "Tokenabcdef****"),
		},
		"IdentifierWithShortToken": {
			id:       types.SessionIdentifier{Token: PString("abc")},
			expected: slog.String("Identifier", "Token????"),
		},
		"IdentifierWithoutIDOrToken": {
			id:       types.SessionIdentifier{},
			expected: slog.String("Identifier", "None Specified"),
		},
	}

	for testName, testCase := range tests {
		t.Run(testName, func(t *testing.T) {
			actualAttr := testCase.id.Slog()
			assert.Equal(t, testCase.expected, actualAttr)
		})
	}
}

func PString(s string) *string {
	return &s
}

func PInt32(i int32) *int32 {
	return &i
}
