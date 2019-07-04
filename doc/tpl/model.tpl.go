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

{{range .models}}
type {{.UpperTableName}} struct{ {{range .Fields}}
    {{.Name}} {{.Type}} {{.Tag}} {{end}}
}
{{if .HasTimeField}}
func (u *{{.UpperTableName}}) MarshalJSON() ([]byte, error) {
	type Alias {{.UpperTableName}}
	t := &struct {
		{{range .TimeFields}} {{.FieldName}} Time {{.FieldTag}}
		{{end}} *Alias
	}{
		{{range .TimeFields}} Time(u.{{.FieldName}}),
		{{end}} (*Alias)(u),
    }
	return json.Marshal(t)
}

func (u *{{.UpperTableName}}) UnmarshalJSON(data []byte) (err error) {
	type Alias {{.UpperTableName}}
	t := &struct {
		{{range .TimeFields}} {{.FieldName}} Time {{.FieldTag}}
		{{end}} *Alias
	}{
		{{range .TimeFields}} Time(u.{{.FieldName}}),
		{{end}} (*Alias)(u),
    }
	err = json.Unmarshal(data, t)
	if err != nil {
		return err
	}
    {{range .TimeFields}} t.Alias.{{.FieldName}} = time.Time(t.{{.FieldName}})
	{{end}}
	*u = {{.UpperTableName}}(*t.Alias)
	return nil
}
{{end}}
{{end}}