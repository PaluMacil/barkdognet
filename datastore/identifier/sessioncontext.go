package identifier

import "github.com/PaluMacil/barkdognet/.gen/barkdog/public/model"

type SessionContext struct {
	User    *model.SysUser
	Roles   []model.SysRole
	Session *model.SysSession
}
