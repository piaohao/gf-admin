package model

import (
	"encoding/json"
	"time"
)

const (
	DateFormat = "2006-01-02"
	TimeFormat = "2006-01-02 15:04:05"
)

type Time time.Time

func Now() Time {
	return Time(time.Now())
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(TimeFormat, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(TimeFormat)
}

type SysPermission struct {
	Id         int       `json:"id" comment:"id"`
	Code       string    `json:"code" comment:"编码"`
	ParentCode string    `json:"parent_code" comment:"父编码"`
	Name       string    `json:"name" comment:"名称"`
	Icon       string    `json:"icon" comment:"图标"`
	Url        string    `json:"url" comment:"链接"`
	Priority   int       `json:"priority" comment:"排序"`
	Level      int       `json:"level" comment:"层级"`
	IsMenu     int       `json:"is_menu" comment:"是否为菜单"`
	Status     string    `json:"status" comment:"状态"`
	CreateTime time.Time `json:"create_time" comment:"创建日期"`
	UpdateTime time.Time `json:"update_time" comment:"更新日期"`
}

func (u *SysPermission) MarshalJSON() ([]byte, error) {
	type Alias SysPermission
	t := &struct {
		CreateTime Time `json:"create_time" comment:"创建日期"`
		UpdateTime Time `json:"update_time" comment:"更新日期"`
		*Alias
	}{
		Time(u.CreateTime),
		Time(u.UpdateTime),
		(*Alias)(u),
	}
	return json.Marshal(t)
}

func (u *SysPermission) UnmarshalJSON(data []byte) (err error) {
	type Alias SysPermission
	t := &struct {
		CreateTime Time `json:"create_time" comment:"创建日期"`
		UpdateTime Time `json:"update_time" comment:"更新日期"`
		*Alias
	}{
		Time(u.CreateTime),
		Time(u.UpdateTime),
		(*Alias)(u),
	}
	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	t.Alias.CreateTime = time.Time(t.CreateTime)
	t.Alias.UpdateTime = time.Time(t.UpdateTime)

	*u = SysPermission(*t.Alias)
	return nil
}

type SysRole struct {
	Id          int       `json:"id" comment:"id"`
	Name        string    `json:"name" comment:"角色名"`
	Description string    `json:"description" comment:"描述"`
	CreateTime  time.Time `json:"create_time" comment:"创建日期"`
	UpdateTime  time.Time `json:"update_time" comment:"更新日期"`
}

func (u *SysRole) MarshalJSON() ([]byte, error) {
	type Alias SysRole
	t := &struct {
		CreateTime Time `json:"create_time" comment:"创建日期"`
		UpdateTime Time `json:"update_time" comment:"更新日期"`
		*Alias
	}{
		Time(u.CreateTime),
		Time(u.UpdateTime),
		(*Alias)(u),
	}
	return json.Marshal(t)
}

func (u *SysRole) UnmarshalJSON(data []byte) (err error) {
	type Alias SysRole
	t := &struct {
		CreateTime Time `json:"create_time" comment:"创建日期"`
		UpdateTime Time `json:"update_time" comment:"更新日期"`
		*Alias
	}{
		Time(u.CreateTime),
		Time(u.UpdateTime),
		(*Alias)(u),
	}
	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	t.Alias.CreateTime = time.Time(t.CreateTime)
	t.Alias.UpdateTime = time.Time(t.UpdateTime)

	*u = SysRole(*t.Alias)
	return nil
}

type SysRolePermission struct {
	Id           int `json:"id" comment:"id"`
	RoleId       int `json:"role_id" comment:"角色id"`
	PermissionId int `json:"permission_id" comment:"权限id"`
}

type SysUser struct {
	Id         int       `json:"id" comment:"用户id"`
	Username   string    `json:"username" comment:"账户名"`
	Password   string    `json:"password" comment:"密码"`
	Salt       string    `json:"salt" comment:"密码盐"`
	Nickname   string    `json:"nickname" comment:"昵称"`
	Mobile     string    `json:"mobile" comment:"电话"`
	Status     string    `json:"status" comment:"状态"`
	CreateTime time.Time `json:"create_time" comment:"创建日期"`
	UpdateTime time.Time `json:"update_time" comment:"更新日期"`
}

func (u *SysUser) MarshalJSON() ([]byte, error) {
	type Alias SysUser
	t := &struct {
		CreateTime Time `json:"create_time" comment:"创建日期"`
		UpdateTime Time `json:"update_time" comment:"更新日期"`
		*Alias
	}{
		Time(u.CreateTime),
		Time(u.UpdateTime),
		(*Alias)(u),
	}
	return json.Marshal(t)
}

func (u *SysUser) UnmarshalJSON(data []byte) (err error) {
	type Alias SysUser
	t := &struct {
		CreateTime Time `json:"create_time" comment:"创建日期"`
		UpdateTime Time `json:"update_time" comment:"更新日期"`
		*Alias
	}{
		Time(u.CreateTime),
		Time(u.UpdateTime),
		(*Alias)(u),
	}
	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}
	t.Alias.CreateTime = time.Time(t.CreateTime)
	t.Alias.UpdateTime = time.Time(t.UpdateTime)

	*u = SysUser(*t.Alias)
	return nil
}

type SysUserRole struct {
	Id     int `json:"id" comment:"id"`
	UserId int `json:"user_id" comment:"用户id"`
	RoleId int `json:"role_id" comment:"角色id"`
}
