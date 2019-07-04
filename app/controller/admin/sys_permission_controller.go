package ctl_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/util/gconv"
	svcAdmin "github.com/piaohao/gf-admin/app/service/admin"
	"github.com/piaohao/gf-admin/util"
	"sort"
)

type SysPermissionController struct {
	gmvc.Controller
}

func (c *SysPermissionController) Index() {
	permissions := c.Session.Get("permissions").(gdb.List)
	util.Html(c.Request, "/admin/layout/layout.html", g.Map{
		"menus":      svcAdmin.SysPermissionService.GetMenu("/admin/sys_permission/index", permissions),
		"contentTpl": "/admin/sys_permission.html",
	})
}

func (c *SysPermissionController) List() {
	list := svcAdmin.SysPermissionService.All()
	util.WriteNormal(c.Request,
		g.Map{
			"recordsTotal":    len(list),
			"recordsFiltered": len(list),
			"data":            list,
		})
}

func (c *SysPermissionController) Tree() {
	allPermissions := svcAdmin.SysPermissionService.All()
	menus := convertToMenu(allPermissions)
	tree := buildTree("super", menus)
	util.WriteNormal(c.Request, tree)
}

func (c *SysPermissionController) RoleTree() {
	allPermissions := svcAdmin.SysPermissionService.All()
	menus := convertToMenu(allPermissions)
	rolePermissions := svcAdmin.SysRolePermissionService.GetByRoleId(c.Request.GetInt("roleId"))
	permissionIds := make([]int, 0)
	for _, rp := range rolePermissions {
		permissionIds = append(permissionIds, rp["permission_id"].(int))
	}
	tree := buildRoleTree("super", menus, permissionIds)
	util.WriteNormal(c.Request, tree)
}

func (c *SysPermissionController) Save() {
	params := c.Request.GetRequestMap()
	if params["id"] == "" {
		svcAdmin.SysPermissionService.Save(gconv.Map(params))
	} else {
		svcAdmin.SysPermissionService.Update(gconv.Map(params))
	}
	util.WriteSuccess(c.Request, nil)
}

func (c *SysPermissionController) Remove() {
	id := c.Request.GetInt("id")
	removeAll(id)
	c.Response.WriteJson(g.Map{"code": 0})
}

func removeAll(id int) {
	permission := svcAdmin.SysPermissionService.Get(id)
	childPermissions := svcAdmin.SysPermissionService.GetByParentCode(gconv.String(permission["code"]))
	if len(childPermissions) == 0 {
		svcAdmin.SysPermissionService.Remove(id)
		return
	}
	for _, r := range childPermissions {
		removeAll(gconv.Int(r["id"]))
	}
}

func buildTree(code string, menus gdb.List) gdb.List {
	var childMenus gdb.List
	for _, menu := range menus {
		if menu["parent_code"] == code {
			childMenus = append(childMenus, menu)
		}
	}
	for _, menu := range childMenus {
		menu["children"] = buildTree(menu["code"].(string), menus)
	}
	ms := svcAdmin.PermissionMap(childMenus)
	sort.Sort(ms)
	return ms
}

func buildRoleTree(code string, menus gdb.List, permissionIds []int) gdb.List {
	var childMenus gdb.List
	for _, menu := range menus {
		if menu["parent_code"] == code {
			childMenus = append(childMenus, menu)
		}
	}
	for _, menu := range childMenus {
		menu["children"] = buildRoleTree(menu["code"].(string), menus, permissionIds)
	}
	for _, menu := range childMenus {
		checked := false
		for _, pId := range permissionIds {
			if menu["id"] == pId {
				checked = true
				break
			}
		}
		menu["checked"] = checked
	}
	ms := svcAdmin.PermissionMap(childMenus)
	sort.Sort(ms)
	return ms
}

func convertToMenu(allPermissions gdb.List) gdb.List {
	var menus gdb.List
	for _, p := range allPermissions {
		isMenuStr := "菜单：是"
		if p["is_menu"] == 0 {
			isMenuStr = "菜单：否"
		}
		m := gdb.Map{
			"id":          p["id"],
			"code":        p["code"],
			"parent_code": p["parent_code"],
			"is_menu":     p["is_menu"],
			"name":        p["name"].(string) + "[" + p["code"].(string) + "]" + "[" + isMenuStr + "]",
			"open":        p["code"] == "root",
			"oName":       p["name"].(string),
			"oUrl":        p["url"].(string),
			"oIcon":       p["icon"].(string),
			"priority":    p["priority"],
		}
		menus = append(menus, m)
	}
	return menus
}
