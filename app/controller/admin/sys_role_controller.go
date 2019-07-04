package ctl_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/util/gconv"
	svcAdmin "github.com/piaohao/gf-admin/app/service/admin"
	"github.com/piaohao/gf-admin/util"
	"strings"
)

type SysRoleController struct {
	gmvc.Controller
}

func (c *SysRoleController) Index() {
	permissions := c.Session.Get("permissions").(gdb.List)
	util.Html(c.Request, "/admin/layout/layout.html", g.Map{
		"menus":      svcAdmin.SysPermissionService.GetMenu("/admin/sys_role/index", permissions),
		"contentTpl": "/admin/sys_role.html",
	})
}

func (c *SysRoleController) List() {
	list := svcAdmin.SysRoleService.All()
	util.WriteNormal(c.Request,
		g.Map{
			"recordsTotal":    len(list),
			"recordsFiltered": len(list),
			"data":            list,
		})
}

func (c *SysRoleController) Get() {
	result := svcAdmin.SysRoleService.Get(c.Request.GetInt("id"))
	util.WriteSuccess(c.Request, result)
}

func (c *SysRoleController) Save() {
	params := c.Request.GetRequestMap()
	permissionIds := strings.Split(params["permissionIds"], ",")
	if params["id"] == "" {
		roleId, _ := svcAdmin.SysRoleService.Save(gconv.Map(params))
		for _, pId := range permissionIds {
			svcAdmin.SysRolePermissionService.Save(g.Map{"role_id": roleId, "permission_id": pId})
		}
	} else {
		svcAdmin.SysRoleService.Update(gconv.Map(params))
		svcAdmin.SysRolePermissionService.RemoveByRoleId(c.Request.GetInt("id"))
		for _, pId := range permissionIds {
			svcAdmin.SysRolePermissionService.Save(g.Map{"role_id": params["id"], "permission_id": pId})
		}
	}
	util.WriteSuccess(c.Request, nil)
}

func (c *SysRoleController) Remove() {
	svcAdmin.SysRoleService.Remove(c.Request.GetInt("id"))
	util.WriteSuccess(c.Request, nil)
}
