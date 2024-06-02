package types_test

import (
	"database/sql/driver"
	"github.com/PaluMacil/barkdognet/datastore/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestTextArray_Scan(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   interface{}
		want    types.TextArray
		wantErr bool
	}{
		{"scan byte array", []byte("{hello,world}"), types.TextArray{"hello", "world"}, false},
		{"scan string", "{hello,world}", types.TextArray{"hello", "world"}, false},
		{"scan empty byte array", []byte("{}"), types.TextArray{}, false},
		{"scan nil", nil, types.TextArray{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ta types.TextArray
			err := ta.Scan(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, ta)
			}
		})
	}
}

func TestTextArray_Value(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   types.TextArray
		want    driver.Value
		wantErr bool
	}{
		{"text array with two values", types.TextArray{"hello", "world"}, `{"hello","world"}`, false},
		{"empty text array", types.TextArray{}, "{}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Value()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestIntArray_Scan(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   interface{}
		want    types.IntArray
		wantErr bool
	}{
		{"scan byte array to int array", []byte("{1,2,3}"), types.IntArray{1, 2, 3}, false},
		{"scan string to int array", "{1,2,3}", types.IntArray{1, 2, 3}, false},
		{"scan empty byte array", []byte("{}"), types.IntArray{}, false},
		{"scan nil", nil, types.IntArray{}, false},
		{"invalid type", true, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ia types.IntArray
			err := ia.Scan(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, ia)
			}
		})
	}
}

func TestIntArray_Value(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		input   types.IntArray
		want    driver.Value
		wantErr bool
	}{
		{"int array to value", types.IntArray{1, 2, 3}, `{1,2,3}`, false},
		{"empty int array", types.IntArray{}, "{}", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.input.Value()
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
