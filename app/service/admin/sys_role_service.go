package svc_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
)

type sysRoleService struct {
}

var SysRoleTable = g.DB().Table("sys_role")
var SysRoleService sysRoleService

func (s *sysRoleService) All() gdb.List {
	results, _ := SysRoleTable.All()
	return results.ToList()
}

func (s *sysRoleService) Get(id int) gdb.Map {
	result, _ := SysRoleTable.Where("id=?", id).One()
	return result.ToMap()
}

func (s *sysRoleService) Save(data g.Map) (int64, error) {
	delete(data, "id")
	result, err := SysRoleTable.Data(data).Filter().Insert()
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *sysRoleService) Update(data g.Map) {
	SysRoleTable.Data(data).Filter().Where("id=?", data["id"]).Update()
}

func (s *sysRoleService) Remove(id int) {
	SysRoleTable.Where("id=?", id).Delete()
}

func (s *sysRoleService) GetByUserId(userId int) gdb.List {
	results, _ := g.DB().GetAll(`select sr.*
	from sys_role sr
	       left join sys_user_role sur on sr.id = sur.role_id
	       left join sys_user su on sur.user_id = su.id
	where su.id = ?`, userId)
	return results.ToList()
}
