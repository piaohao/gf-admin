package svc_admin

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/util/gconv"
	"github.com/piaohao/gf-admin/util"
)

type sysUserService struct {
}

var SysUserTable = g.DB().Table("sys_user")
var SysUserService sysUserService

func (s *sysUserService) Table() *gdb.Model {
	return SysUserTable
}

func (s *sysUserService) Page(page, size int) gdb.List {
	results, _ := SysUserTable.ForPage(page+1, size).All()
	return results.ToList()
}

func (s *sysUserService) Count() int {
	result, _ := SysUserTable.Count()
	return result
}

func (s *sysUserService) All() gdb.List {
	results, _ := SysUserTable.All()
	return results.ToList()
}

func (s *sysUserService) Get(id int) gdb.Map {
	result, _ := SysUserTable.Where("id=?", id).One()
	return result.ToMap()
}

func (s *sysUserService) Save(data g.Map) (int64, error) {
	delete(data, "id")
	result, err := SysUserTable.Data(data).Filter().Insert()
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *sysUserService) Update(data g.Map) {
	SysUserTable.Data(data).Filter().Where("id=?", data["id"]).Update()
}

func (s *sysUserService) Remove(id int) {
	result, _ := SysUserTable.Where("id=?", id).One()
	if result != nil {
		user := result.ToMap()
		if user["name"] != "admin" {
			SysUserTable.Where("id=?", id).Delete()
		}
	}
}

//====custom method====//
func (s *sysUserService) GetByLoginInfo(username, password string) gdb.Map {
	result, _ := SysUserTable.Where("username=?", username).One()
	if result == nil {
		return nil
	}
	user := result.ToMap()
	if util.Md5WithSalt(password, gconv.String(user["salt"])) != gconv.String(user["password"]) {
		return nil
	}
	return user
}
