//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type SysSession struct {
	ID           int32 `sql:"primary_key"`
	SysUserID    int32
	SessionToken string
	CreatedAt    time.Time
}