package {{.CtlPkgName}}

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/util/gconv"
	"{{.ModuleName}}/app/service/admin"
	"{{.ModuleName}}/util"
)

type {{.UpperCamelTableName}}Controller struct {
	gmvc.Controller
}

func (c *{{.UpperCamelTableName}}Controller) Index() {
	permissions := c.Session.Get("permissions").(gdb.List)
	util.Html(c.Request, "/admin/layout/layout.html", g.Map{
		"menus":      svc_admin.SysPermissionService.GetMenu("/admin/{{.TableName}}/index", permissions),
		"contentTpl": "/admin/{{.TableName}}.html",
	})
}

func (c *{{.UpperCamelTableName}}Controller) List() {
	user := c.Session.Get("user").(gdb.Map)
	params := c.Request.GetMap()
	if params["proxy_id"] == "" && !util.IsAdmin(c.Session) {
		params["proxy_id"] = gconv.String(user["id"])
	}
	list,total := {{.SvcPkgName}}.{{.UpperCamelTableName}}Service.Page(params)
	util.WriteNormal(c.Request,
		g.Map{
			"recordsTotal":    total,
			"recordsFiltered": total,
			"data":            list,
		})
}

func (c *{{.UpperCamelTableName}}Controller) Get() {
	result := {{.SvcPkgName}}.{{.UpperCamelTableName}}Service.Get(c.Request.GetInt("id"))
	util.WriteSuccess(c.Request, result)
}

func (c *{{.UpperCamelTableName}}Controller) Save() {
	params := c.Request.GetRequestMap()
	if params["id"] == "" {
		{{.SvcPkgName}}.{{.UpperCamelTableName}}Service.Save(gconv.Map(params))
	} else {
		{{.SvcPkgName}}.{{.UpperCamelTableName}}Service.Update(gconv.Map(params))
	}
	util.WriteSuccess(c.Request, nil)
}

func (c *{{.UpperCamelTableName}}Controller) Remove() {
	{{.SvcPkgName}}.{{.UpperCamelTableName}}Service.Remove(c.Request.GetInt("id"))
	util.WriteSuccess(c.Request, nil)
}
