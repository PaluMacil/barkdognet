package aggregates

import "github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"

type SessionDetails struct {
	User    *model.SysUser
	Roles   []model.SysRole
	Session *model.SysSession
}
