package ctl_admin

import (
	"bufio"
	"bytes"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/frame/gmvc"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/text/gstr"
	svcAdmin "github.com/piaohao/gf-admin/app/service/admin"
	"github.com/piaohao/gf-admin/util"
	"html/template"
	"os"
	"strings"
)

type CodeController struct {
	gmvc.Controller
}

type field_ struct {
	OriginalName template.JS
	Name         string
	Type         string
	Tag          template.HTML
	Comment      string
}

type timeField_ struct {
	FieldName string
	FieldTag  template.HTML
}

type model_ struct {
	TableName      string
	UpperTableName string
	Fields         []field_
	HasTimeField   bool
	TimeFields     []timeField_
}

func (c *CodeController) Index() {
	permissions := c.Session.Get("permissions").(gdb.List)
	util.Html(c.Request, "/admin/layout/layout.html", g.Map{
		"menus":      svcAdmin.SysPermissionService.GetMenu("/admin/code/index", permissions),
		"contentTpl": "/admin/sys_code.html",
	})
}

func (c *CodeController) Gen() {
	tableName := c.Request.GetString("table")
	svcPkgName := c.Request.GetString("svcPkgName")
	ctlPkgName := c.Request.GetString("ctlPkgName")
	if svcPkgName == "" {
		svcPkgName = "svc_admin_biz"
	}
	if ctlPkgName == "" {
		ctlPkgName = "ctl_admin_biz"
	}
	tables := make([]string, 0)
	if tableName == "" {
		tableResult, _ := g.DB().GetAll(`select table_name
	from information_schema.TABLES
	where table_schema = ?`, g.Config().GetString("code.database"))
		for _, record := range tableResult.ToList() {
			tables = append(tables, record["table_name"].(string))
		}
	} else {
		tables = append(tables, tableName)
	}
	generate(tables, svcPkgName, ctlPkgName)
	c.Response.Write("生成成功！")
}

func generate(tables []string, svcPkgName, ctlPkgName string) {
	moduleName := g.Config().GetString("code.modulename")
	models := make([]model_, 0)
	os.RemoveAll("generateCode")
	os.MkdirAll("generateCode", os.ModeDir)
	for _, tableName := range tables {
		result, _ := g.DB().GetAll(`
		select column_name,data_type,column_key,column_comment from information_schema.columns
		where table_schema = ? 
		and table_name = ?
		`, g.Config().GetString("code.database"), tableName)
		tableCommentResult, _ := g.DB().GetValue(`
		SELECT t.TABLE_COMMENT
		FROM information_schema.TABLES t
		WHERE t.TABLE_SCHEMA = ?
		  and t.TABLE_NAME = ?
		`, g.Config().GetString("code.database"), tableName)
		tableComment := tableCommentResult.String()
		fields := make([]field_, 0)
		timeFields := make([]timeField_, 0)
		for _, r := range result.ToList() {
			f := field_{}
			f.OriginalName = template.JS(r["column_name"].(string))
			f.Name = convertField(r["column_name"].(string))
			f.Type = convertType(r["data_type"].(string))
			f.Tag = convertTag(r["column_name"].(string), r["column_comment"].(string))
			f.Comment = r["column_comment"].(string)
			fields = append(fields, f)
			if f.Type == "time.Time" {
				timeFields = append(timeFields, timeField_{FieldName: f.Name, FieldTag: f.Tag})
			}
		}
		upperCamelTableName := convertField(tableName)
		camelTableName := gstr.LcFirst(upperCamelTableName)

		{
			tpl := template.Must(template.New("html.tpl.html").ParseFiles("doc/tpl/html.tpl.html"))
			var filename = "generateCode/" + tableName + ".html"
			f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
			if err != nil {
				println(err.Error())
			}
			bw := bufio.NewWriter(f)
			err = tpl.Execute(bw, g.Map{
				"Fields":       fields,
				"TableComment": tableComment,
				"TableName":    tableName,
			})
			if err != nil {
				glog.Errorfln("html解析错误:%v", err)
			}
			bw.Flush()
			f.Close()
		}
		{
			tpl := template.Must(template.New("service.tpl.go").ParseFiles("doc/tpl/service.tpl.go"))
			var filename = "generateCode/" + tableName + "_service.go"
			f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
			if err != nil {
				println(err.Error())
			}
			bw := bufio.NewWriter(f)
			err = tpl.Execute(bw, g.Map{
				"Fields":              fields,
				"TableComment":        tableComment,
				"TableName":           tableName,
				"CamelTableName":      camelTableName,
				"UpperCamelTableName": upperCamelTableName,
				"SvcPkgName":          svcPkgName,
				"CtlPkgName":          ctlPkgName,
				"ModuleName":          moduleName,
			})
			if err != nil {
				glog.Errorfln("service解析错误:%v", err)
			}
			bw.Flush()
			f.Close()
		}
		{
			tpl := template.Must(template.New("controller.tpl.go").ParseFiles("doc/tpl/controller.tpl.go"))
			var filename = "generateCode/" + tableName + "_controller.go"
			f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
			if err != nil {
				println(err.Error())
			}
			bw := bufio.NewWriter(f)
			err = tpl.Execute(bw, g.Map{
				"Fields":              fields,
				"TableComment":        tableComment,
				"TableName":           tableName,
				"CamelTableName":      camelTableName,
				"UpperCamelTableName": upperCamelTableName,
				"SvcPkgName":          svcPkgName,
				"CtlPkgName":          ctlPkgName,
				"ModuleName":          moduleName,
			})
			if err != nil {
				glog.Errorfln("controller解析错误:%v", err)
			}
			bw.Flush()
			f.Close()
		}

		model := model_{}
		model.TableName = camelTableName
		model.UpperTableName = gstr.UcFirst(camelTableName)
		model.Fields = fields
		if len(timeFields) > 0 {
			model.HasTimeField = true
			model.TimeFields = timeFields
		}
		models = append(models, model)
	}
	{
		var filename = "generateCode/model.go"
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666) //打开文件
		if err != nil {
			println(err.Error())
		}
		tpl := template.Must(template.New("model.tpl.go").ParseFiles("doc/tpl/model.tpl.go"))
		bw := bufio.NewWriter(f)
		err = tpl.Execute(bw, g.Map{
			"models": models,
		})
		if err != nil {
			glog.Errorfln("model解析错误:%v", err)
		}
		bw.Flush()
		f.Close()
	}
}

func convertTag(columnName, columnComment string) template.HTML {
	tagTpl := "`json:\"{{.jsonName}}\" comment:\"{{.comment}}\"`"
	tpl := template.Must(template.New("").Parse(tagTpl))
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	tpl.Execute(bw, g.Map{
		"jsonName": columnName,
		"comment":  columnComment,
	})
	bw.Flush()
	return template.HTML(string(b.Bytes()))
}

func convertType(columnType string) string {
	switch columnType {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		return "byte[]"
	case "bit", "int", "tinyint", "small_int", "medium_int":
		return "int"
	case "big_int":
		return "int64"
	case "float", "double", "decimal":
		return "float64"
	case "bool":
		return "bool"
	case "timestamp", "date", "datetime":
		return "time.Time"
	default:
		// 自动识别类型, 以便默认支持更多数据库类型
		switch {
		case strings.Contains(columnType, "int"):
			return "int"
		case strings.Contains(columnType, "text") || strings.Contains(columnType, "char"):
			return "string"
		case strings.Contains(columnType, "float") || strings.Contains(columnType, "double"):
			return "float64"
		case strings.Contains(columnType, "bool"):
			return "bool"
		case strings.Contains(columnType, "binary") || strings.Contains(columnType, "blob"):
			return "byte[]"
		default:
			return "string"
		}
	}
}

func convertField(columnName string) string {
	arr := strings.Split(columnName, "_")
	newArr := make([]string, 0)
	for _, s := range arr {
		newArr = append(newArr, gstr.UcFirst(s))
	}
	return strings.Join(newArr, "")
}
