package ctl_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/util/gconv"
	svcAdmin "github.com/piaohao/gf-admin/app/service/admin"
	"github.com/piaohao/gf-admin/util"
)

type IndexController struct {
	gmvc.Controller
}

func (c *IndexController) Index() {
	c.Response.Write("welcome to gf-admin")
}

func (c *IndexController) DoLogin() {
	r := c.Request
	username := r.GetPostString("username")
	password := r.GetPostString("password")
	user := svcAdmin.SysUserService.GetByLoginInfo(username, password)
	if user == nil {
		util.Html(c.Request, "/admin/sys_login.html", g.Map{"error": "用户名或密码错误！"})
	} else {
		r.Session.Set("user", user)
		r.Session.Set("roles", svcAdmin.SysRoleService.GetByUserId(gconv.Int(user["id"])))
		r.Session.Set("permissions", svcAdmin.SysPermissionService.GetByUserId(gconv.Int(user["id"])))
		r.Response.RedirectTo("/admin/main")
	}
}

func (c *IndexController) DoLogout() {
	r := c.Request
	r.Session.Clear()
	util.Html(c.Request, "/admin/sys_login.html")
}

func (c *IndexController) Login() {
	//c.Session.Id()
	util.Html(c.Request, "/admin/sys_login.html")
}

func (c *IndexController) Main() {
	permissions := c.Session.Get("permissions").(gdb.List)
	util.Html(c.Request, "/admin/layout/layout.html", g.Map{
		"menus":      svcAdmin.SysPermissionService.GetMenu("/admin/main", permissions),
		"contentTpl": "/admin/sys_main.html",
	})
}
