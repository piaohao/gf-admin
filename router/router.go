package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/encoding/gjson"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
	ctlAdmin "github.com/piaohao/gf-admin/app/controller/admin"
	ctlWeb "github.com/piaohao/gf-admin/app/controller/web"
	"github.com/piaohao/gf-admin/util"
	"strings"
)

// 统一路由注册.
func init() {
	{
		adminGroup := g.Server().Group("/admin")
		adminGroup.Bind([]ghttp.GroupItem{
			{"ALL", "*", AdminHookHandler, ghttp.HOOK_BEFORE_SERVE},
			{"ALL", "/", new(ctlAdmin.IndexController)},
			{"ALL", "/code", new(ctlAdmin.CodeController)},
			{"ALL", "/sys_user", new(ctlAdmin.SysUserController)},
			{"ALL", "/sys_role", new(ctlAdmin.SysRoleController)},
			{"ALL", "/sys_permission", new(ctlAdmin.SysPermissionController)},
		})
	}
	{
		webGroup := g.Server().Group("/web")
		webGroup.Bind([]ghttp.GroupItem{
			{"ALL", "*", WebHookHandler, ghttp.HOOK_BEFORE_SERVE},
			{"ALL", "/my", new(ctlWeb.MyController)},
		})
	}
	g.Server().SetLogHandler(func(r *ghttp.Request, err ...interface{}) {
		jsonArr, _ := gjson.Encode(r.GetMap())
		content := fmt.Sprintf(`路径:%s,方法:%s,状态:%d,耗时:%d毫秒,参数:%s,客户端IP:%s,服务器地址:%s`,
			r.URL.RequestURI(),
			r.Method,
			r.Response.Writer.Status,
			gconv.Int((r.LeaveTime-r.EnterTime)/1000),
			string(jsonArr),
			r.GetClientIp(),
			r.Host)
		if err != nil {
			glog.Cat("error").Backtrace(true, 2).Error(err)
			requestType := r.Header.Get("X-Requested-With")
			r.Response.ClearBuffer()
			if "XMLHttpRequest" == requestType {
				r.Response.WriteJson(g.Map{"code": 1, "message": "内部错误"})
			} else {
				r.Response.Write("系统内部错误")
			}
			r.Response.Output()
		} else {
			glog.Cat("access").Backtrace(false, 2).Println(content)
		}
	})
}

var adminAnonUrls = []string{"/admin/index", "/admin/login", "/admin/doLogin", "/admin/code/gen"}

func AdminHookHandler(r *ghttp.Request) {
	for _, url := range adminAnonUrls {
		if url == r.Request.URL.Path {
			return
		}
	}
	if r.Session.Get("user") == nil {
		requestType := r.Header.Get("X-Requested-With")
		if "XMLHttpRequest" == requestType {
			util.WriteError(r, 10000, "用户登录信息已过期！")
		} else {
			util.Html(r, "/admin/sys_login.html", g.Map{"error": "用户登录信息已过期！"})
		}
		r.ExitAll()
	}
}

var webAnonUrls = []string{"/web/my/auth"}

func WebHookHandler(r *ghttp.Request) {
	for _, url := range webAnonUrls {
		if url == r.Request.URL.Path {
			return
		}
	}
	requestType := r.Header.Get("X-Requested-With")
	if "XMLHttpRequest" == requestType {
		requestHeader := r.Header.Get("Authorization")
		authToken := ""
		if requestHeader != "" && strings.HasPrefix(requestHeader, "Bearer ") {
			authToken = requestHeader[7:]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				return []byte("gf-cli"), nil
			})
			if err != nil {
				util.WriteError(r, 700, "登录信息已过期,请重新登录")
			}
			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				util.WriteError(r, 700, "登录信息已过期,请重新登录")
			}
			//验证token，如果token被修改过则为false
			if !token.Valid {
				util.WriteError(r, 700, "登录信息已过期,请重新登录")
			}

			r.SetParam("userId", claim["sub"])
		}
	}
}
