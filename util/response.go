package util

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/util/gconv"
)

func WriteError(r *ghttp.Request, code int, message string) {
	if err := r.Response.WriteJson(g.Map{"code": code, "message": message}); err != nil {
		glog.Errorf("出错了:%s", message)
	}
	r.ExitAll()
}

func WriteErrorByDefaultCode(r *ghttp.Request, message string) {
	if err := r.Response.WriteJson(g.Map{"code": 1, "message": message}); err != nil {
		glog.Errorf("出错了:%s", message)
	}
	r.ExitAll()
}

func WriteDefaultError(r *ghttp.Request) {
	if err := r.Response.WriteJson(g.Map{"code": 1, "message": "内部错误"}); err != nil {
		glog.Errorf("输出json出错")
	}
	r.ExitAll()
}

func WriteSuccess(r *ghttp.Request, data interface{}) {
	if err := r.Response.WriteJson(g.Map{"code": 0, "message": "success", "data": data}); err != nil {
		glog.Errorf("输出json出错")
	}
	r.ExitAll()
}

func WriteNormal(r *ghttp.Request, data interface{}) {
	if err := r.Response.WriteJson(data); err != nil {
		glog.Errorf("输出json出错")
	}
	r.ExitAll()
}

func Html(r *ghttp.Request, view string, data ...interface{}) {
	params := make(g.Map)
	funcs := make(g.Map)
	if len(data) == 1 {
		params = gconv.Map(data[0])
	} else if len(data) == 2 {
		params = gconv.Map(data[0])
		funcs = gconv.Map(data[1])
	} else if len(data) > 2 {
		params = gconv.Map(data[0])
		funcs = gconv.Map(data[1])
	}
	err := r.Response.WriteTpl(view, params, funcs)
	if err != nil {
		panic(err)
	}
}
