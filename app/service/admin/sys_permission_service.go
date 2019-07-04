package svc_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/util/gconv"
	"sort"
)

type sysPermissionService struct {
}

var SysPermissionTable = g.DB().Table("sys_permission")
var SysPermissionService sysPermissionService

func (s *sysPermissionService) All() gdb.List {
	results, _ := SysPermissionTable.Cache(0, "allPermissions").All()
	return results.ToList()
}

func (s *sysPermissionService) Get(id int) gdb.Map {
	result, _ := SysPermissionTable.Where("id=?", id).One()
	return result.ToMap()
}

func (s *sysPermissionService) Save(data g.Map) (int64, error) {
	delete(data, "id")
	data["priority"] = gconv.Int(data["priority"])
	result, err := SysPermissionTable.Cache(-1, "allPermissions").Data(data).Filter().Insert()
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *sysPermissionService) Update(data g.Map) {
	data["priority"] = gconv.Int(data["priority"])
	SysPermissionTable.Cache(-1, "allPermissions").Data(data).Filter().Where("id=?", data["id"]).Update()
}

func (s *sysPermissionService) Remove(id int) {
	SysPermissionTable.Cache(-1, "allPermissions").Where("id=?", id).Delete()
}

//====custom method====//

func (s *sysPermissionService) GetByParentCode(parentCode string) gdb.List {
	results, _ := SysPermissionTable.Where("parent_code=?", parentCode).All()
	return results.ToList()
}

func (s *sysPermissionService) GetMenu(url string, permissions gdb.List) string {
	results, _ := SysPermissionTable.Cache(0, "allMenus").Where("status='Y' and is_menu=1").Select()
	parentCode := s.getParentCode(url, results.ToList())
	menus := s.buildMenu("root", results.ToList(), permissions)
	var selectMain = false
	if parentCode == "" {
		selectMain = true
	}
	customTpl := `
		<li class="nav-item {{if eq (.parentCode|len) 0 }}{{if eq .viewUrl "/admin/main"}}active{{end}}{{end}}">
          <a class="nav-link" href="/admin/main">
            <i class="fas fa-fw fa-tachometer-alt"></i>
            <span>主页</span>
          </a>
        </li>
		{{$parentCode:=.parentCode}}
		{{$viewUrl:=.viewUrl}}
		{{range .menus}}
		<li class="nav-item {{if gt (.children|len) 0 }}dropdown{{end}} {{if eq $parentCode .code}}active show{{end}}{{if eq ($parentCode|len) 0 }}{{if eq $viewUrl .url}}active{{end}}{{end}}">
          <a class="nav-link {{if gt (.children|len) 0 }}dropdown-toggle{{end}}" href="{{.url}}" {{if gt (.children|len) 0 }} role="button" data-toggle="dropdown" aria-haspopup="true"{{end}} aria-expanded=" {{if eq $parentCode .code}} true{{end}}">
            {{if ne .icon ""}}
              <i class="fas fa-fw {{.icon}}"></i>
		    {{end}}
            <span>{{.name}}</span>
          </a>
          {{if gt (.children|len) 0 }}
          <div class="dropdown-menu  {{if eq $parentCode .code}} show{{end}}" aria-labelledby="pagesDropdown" style="">
            {{range .children}}
            <a class="dropdown-item {{if eq $viewUrl .url}} active{{end}}" href="{{.url}}">
              {{if ne .icon ""}}
              <i class="fas fa-fw {{.icon}}"></i>
		      {{end}}
              <span>{{.name}}</span>
            </a>
            {{end}}
          </div>
          {{end}}
        </li>
		{{end}}
	`
	bytes, err := g.View().ParseContent(customTpl, g.Map{"menus": menus, "parentCode": parentCode, "viewUrl": url, "selectMain": selectMain})
	if err == nil {
		return string(bytes)
	} else {
		println(err.Error())
		return ""
	}
}

type PermissionMap gdb.List

// 获取此 slice 的长度
func (p PermissionMap) Len() int { return len(p) }

// 比较两个元素大小 升序
func (p PermissionMap) Less(i, j int) bool { return p[i]["priority"].(int) < p[j]["priority"].(int) }

// 交换数据
func (p PermissionMap) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

func (s *sysPermissionService) buildMenu(code string, menus gdb.List, permissions gdb.List) gdb.List {
	var childMenus gdb.List
	for _, menu := range menus {
		if menu["parent_code"] == code {
			exist := false
			for _, p := range permissions {
				if menu["id"] == p["id"] {
					exist = true
					break
				}
			}
			if exist {
				childMenus = append(childMenus, menu)
			}
		}
	}
	for _, menu := range childMenus {
		menu["children"] = s.buildMenu(menu["code"].(string), menus, permissions)
	}
	ms := PermissionMap(childMenus)
	sort.Sort(ms)
	return ms
}

func (s *sysPermissionService) getParentCode(url string, menus gdb.List) string {
	parentCode := ""
	for _, menu := range menus {
		if menu["url"] == url {
			parentCode = menu["parent_code"].(string)
			break
		}
	}
	for _, menu := range menus {
		if menu["code"] == parentCode {
			return menu["code"].(string)
		}
	}
	return ""
}

func (s *sysPermissionService) GetByUserId(userId int) gdb.List {
	results, _ := g.DB().GetAll(`
		select sp.*
		from sys_permission sp
		       left join sys_role_permission srp on sp.id = srp.permission_id
		       left join (select sr.*, su.id user_id
		                  from sys_role sr
		                         right join sys_user_role sur on sr.id = sur.role_id
		                         right join sys_user su on su.id = sur.user_id
		) a on a.id = srp.role_id
		where a.user_id = ?;
		`, userId)
	return results.ToList()
}

func (s *sysPermissionService) GetPermissionsV2(userId int) gdb.List {
	permissions := s.GetByUserId(userId)
	return s.wrapPermissionsV2("root", permissions)
}

func (s *sysPermissionService) wrapPermissionsV2(code string, permissions gdb.List) gdb.List {
	var childMenus gdb.List
	for _, menu := range permissions {
		if menu["parent_code"] == code && menu["is_menu"] == 1 {
			item := make(g.Map)
			item["path"] = menu["url"]
			item["name"] = menu["code"]
			item["component"] = menu["view"]
			item["meta"] = g.Map{
				"title": menu["name"],
				"icon":  menu["icon"],
			}
			item["priority"] = menu["priority"]

			childMenus = append(childMenus, item)
		}
	}
	for _, menu := range childMenus {
		menu["children"] = s.wrapPermissionsV2(menu["name"].(string), permissions)
	}
	ms := PermissionMap(childMenus)
	sort.Sort(ms)
	return ms
}
