package svc_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
)

type sysRolePermissionService struct {
}

var SysRolePermissionTable = g.DB().Table("sys_role_permission")
var SysRolePermissionService sysRolePermissionService

func (s *sysRolePermissionService) All() gdb.List {
	results, _ := SysRolePermissionTable.All()
	return results.ToList()
}

func (s *sysRolePermissionService) Get(id int) gdb.Map {
	result, _ := SysRolePermissionTable.Where("id=?", id).One()
	return result.ToMap()
}

func (s *sysRolePermissionService) Save(data g.Map) (int64, error) {
	delete(data, "id")
	result, err := SysRolePermissionTable.Data(data).Filter().Insert()
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *sysRolePermissionService) Update(data g.Map) {
	SysRolePermissionTable.Data(data).Filter().Where("id=?", data["id"]).Update()
}

func (s *sysRolePermissionService) Remove(id int) {
	SysRolePermissionTable.Where("id=?", id).Delete()
}

func (s *sysRolePermissionService) GetByRoleId(roleId int) gdb.List {
	results, _ := SysRolePermissionTable.Where("role_id=?", roleId).All()
	return results.ToList()
}

func (s *sysRolePermissionService) RemoveByRoleId(roleId int) {
	SysRolePermissionTable.Where("role_id=?", roleId).Delete()
}
