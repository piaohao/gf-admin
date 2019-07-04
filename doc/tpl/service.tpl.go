package {{.SvcPkgName}}

import (
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/database/gdb"
	"github.com/gogf/gf/g/util/gconv"
)

type {{.CamelTableName}}Service struct {
}

var {{.UpperCamelTableName}}Table = g.DB().Table("{{.TableName}}")
var {{.UpperCamelTableName}}Service {{.CamelTableName}}Service

func (s *{{.CamelTableName}}Service) Page(params map[string]string) (gdb.List,int) {
	table := {{.CamelTableName}}Table
	pageTable := table.Clone()
	table = table.Where("1=1")
	pageTable = pageTable.Where("1=1")
	{{range .Fields}}
	if params["{{.OriginalName}}"] != "" {
		table = table.And("{{.OriginalName}}=?", params["{{.OriginalName}}"])
		pageTable = pageTable.And("{{.OriginalName}}=?", params["{{.OriginalName}}"])
	}
	{{end}}
	results, err := table.ForPage(gconv.Int(params["page"]), gconv.Int(params["pageSize"])).All()
	if err != nil {
		panic(err)
	}
	total, err := pageTable.Count()
	if err != nil {
		panic(err)
	}
	return results.ToList(),total
}

func (s *{{.CamelTableName}}Service) All() gdb.List {
	results, err := {{.CamelTableName}}Table.All()
	if err != nil {
		panic(err)
	}
	return results.ToList()
}

func (s *{{.CamelTableName}}Service) Get(id int) gdb.Map {
	result, err := {{.CamelTableName}}Table.Where("id=?", id).One()
	if err != nil {
		panic(err)
	}
	return result.ToMap()
}

func (s *{{.CamelTableName}}Service) Save(data g.Map) (int64, error) {
	delete(data, "id")
	result, err := {{.CamelTableName}}Table.Data(data).Filter().Insert()
	if err != nil {
		return -1, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (s *{{.CamelTableName}}Service) Update(data g.Map) {
	_ , err := {{.CamelTableName}}Table.Data(data).Filter().Where("id=?", data["id"]).Update()
	if err != nil {
		panic(err)
	}
}

func (s *{{.CamelTableName}}Service) Remove(id int) {
	_ , err := {{.CamelTableName}}Table.Where("id=?", id).Delete()
	if err != nil {
		panic(err)
	}
}