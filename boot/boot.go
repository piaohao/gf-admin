package boot

import (
	"flag"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/gcfg"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/os/gview"
)

// 用于应用初始化。
func init() {
	v := g.View()
	c := g.Config()
	s := g.Server()

	initConfig(c)
	initView(v, c)

	debugMode := c.GetBool("debug")
	// glog配置
	logpath := c.GetString("logpath")
	glog.SetPath(logpath)
	glog.SetStdoutPrint(true)
	if debugMode {
		glog.SetLevel(glog.LEVEL_ALL)
	} else {
		glog.SetLevel(glog.LEVEL_INFO | glog.LEVEL_ERRO | glog.LEVEL_CRIT)
	}

	// db配置
	g.DB().SetDebug(debugMode)

	// Web Server配置
	s.EnableAdmin()
	s.EnablePprof()

	s.BindStatusHandler(404, func(r *ghttp.Request) {
		r.Response.Write("404 not found!")
	})
	s.SetServerRoot("public")
	s.SetLogPath(logpath)
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_CAMEL)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetDumpRouteMap(debugMode)
	port := c.GetInt("server.port")
	if port == 0 {
		port = 8080
	}
	s.SetPort(port)
}

func initConfig(c *gcfg.Config) {
	_ = c.AddPath("config")
	var configFileName string
	flag.StringVar(&configFileName, "config.file.name", "config-dev.yml", "config.file.name")
	flag.Parse()
	c.SetFileName(configFileName)
}

func initView(v *gview.View, c *gcfg.Config) {
	_ = v.AddPath("template")
	v.SetDelimiters("{{", "}}")
	v.BindFunc("hasPermission", func(permission string, allPermissions gdb.List) bool {
		for _, p := range allPermissions {
			if permission == p["url"] {
				return true
			}
		}
		return false
	})
	v.BindFunc("isAdmin", func(session g.Map) bool {
		for _, r := range session["roles"].(gdb.List) {
			if "admin" == r["name"] {
				return true
			}
		}
		return false
	})
	v.Assign("OssUrlPrefix", c.GetString("oss.url.prefix"))
}
