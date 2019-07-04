package svc_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
)

type sysUserRoleService struct {
}

var SysUserRoleTable = g.DB().Table("sys_user_role")
var SysUserRoleService sysUserRoleService

func (s *sysUserRoleService) All() gdb.List {
	results, _ := SysUserRoleTable.All()
	return results.ToList()
}

func (s *sysUserRoleService) Get(id int) gdb.Map {
	result, _ := SysUserRoleTable.Where("id=?", id).One()
	return result.ToMap()
}

func (s *sysUserRoleService) Save(data g.Map) (int64, error) {
	delete(data, "id")
	result, err := SysUserRoleTable.Data(data).Filter().Insert()
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *sysUserRoleService) Update(data g.Map) {
	SysUserRoleTable.Data(data).Filter().Where("id=?", data["id"]).Update()
}

func (s *sysUserRoleService) Remove(id int) {
	SysUserRoleTable.Where("id=?", id).Delete()
}

func (s *sysUserRoleService) GetByUserId(userId int) gdb.List {
	results, _ := SysUserRoleTable.Where("user_id=?", userId).All()
	return results.ToList()
}

func (s *sysUserRoleService) RemoveByUserId(userId int) {
	SysUserRoleTable.Where("user_id=?", userId).Delete()
}
