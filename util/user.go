package util

import (
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/net/ghttp"
)

func HasPermission(permission string, allPermissions gdb.List) bool {
	for _, p := range allPermissions {
		if permission == p["url"] {
			return true
		}
	}
	return false
}

func IsAdmin(session *ghttp.Session) bool {
	for _, r := range session.Get("roles").(gdb.List) {
		if "admin" == r["name"] {
			return true
		}
	}
	return false
}
