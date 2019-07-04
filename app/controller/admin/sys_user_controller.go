package ctl_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/gogf/gf/g/util/grand"
	svcAdmin "github.com/piaohao/gf-admin/app/service/admin"
	"github.com/piaohao/gf-admin/util"
	"strings"
)

type SysUserController struct {
	gmvc.Controller
}

func (c *SysUserController) Index() {
	permissions := c.Session.Get("permissions").(gdb.List)
	util.Html(c.Request, "/admin/layout/layout.html", g.Map{
		"menus":      svcAdmin.SysPermissionService.GetMenu("/admin/sys_user/index", permissions),
		"contentTpl": "/admin/sys_user.html",
	})
}

func (c *SysUserController) List() {
	params := c.Request.GetMap()
	table := g.DB().Table("sys_user su")
	pageTable := table.Clone()
	table = table.LeftJoin("sys_user_role sur", "sur.user_id=su.id").LeftJoin("sys_role sr", "sur.role_id=sr.id")
	table = table.Where("1=1")
	pageTable = pageTable.Where("1=1")
	if params["username"] != "" {
		table = table.And("username=?", params["username"])
		pageTable = pageTable.And("username=?", params["username"])
	}
	if params["nickname"] != "" {
		table = table.And("nickname=?", params["nickname"])
		pageTable = pageTable.And("nickname=?", params["nickname"])
	}
	results, _ := table.ForPage(gconv.Int(params["page"]), gconv.Int(params["pageSize"])).Fields("su.*,sr.name role_name").All()
	list := results.ToList()
	total, _ := pageTable.Count()
	util.WriteNormal(c.Request,
		g.Map{
			"recordsTotal":    total,
			"recordsFiltered": total,
			"data":            list,
		})
}

func (c *SysUserController) Get() {
	result := svcAdmin.SysUserService.Get(c.Request.GetInt("id"))
	if result == nil {
		util.WriteErrorByDefaultCode(c.Request, "记录不存在")
	}
	util.WriteSuccess(c.Request, result)
}

func (c *SysUserController) GetRoleIds() {
	userRoles := svcAdmin.SysUserRoleService.GetByUserId(c.Request.GetInt("id"))
	allRoles := svcAdmin.SysRoleService.All()
	util.WriteSuccess(c.Request, g.Map{"userRoles": userRoles, "allRoles": allRoles})
}

func (c *SysUserController) Save() {
	params := c.Request.GetRequestMap()
	roleIds := strings.Split(params["roleIds"], ",")
	if params["id"] == "" {
		params["salt"] = grand.RandStr(6)
		params["password"] = util.Md5WithSalt(params["password"], params["salt"])
		user := gconv.Map(params)
		if user == nil {
			glog.Errorfln("结构体转换错误")
			util.WriteDefaultError(c.Request)
		}
		userId, err := svcAdmin.SysUserService.Save(user)
		if err != nil {
			glog.Errorfln("保存用户错误:%v", err)
			util.WriteDefaultError(c.Request)
		}
		for _, rId := range roleIds {
			svcAdmin.SysUserRoleService.Save(g.Map{"role_id": rId, "user_id": userId})
		}
	} else {
		record := svcAdmin.SysUserService.Get(c.Request.GetInt("id"))
		if params["password"] != record["password"].(string) {
			params["password"] = util.Md5WithSalt(params["password"], record["salt"].(string))
		}
		user := gconv.Map(params)
		if user == nil {
			glog.Errorfln("结构体转换错误")
			util.WriteDefaultError(c.Request)
		}
		svcAdmin.SysUserService.Update(user)
		svcAdmin.SysUserRoleService.RemoveByUserId(c.Request.GetInt("id"))
		for _, rId := range roleIds {
			svcAdmin.SysUserRoleService.Save(g.Map{"role_id": rId, "user_id": params["id"]})
		}
	}
	util.WriteSuccess(c.Request, nil)
}

func (c *SysUserController) Remove() {
	svcAdmin.SysUserService.Remove(c.Request.GetInt("id"))
	util.WriteSuccess(c.Request, nil)
}
