package printTable

import (
	"fmt"
	"reflect"
)

type PrintTable struct {
	header []string
	layout string
}

func NewPrintTable(header []string, layout string) *PrintTable {
	return &PrintTable{
		header: header,
		layout: layout,
	}
}

// Format retorna a table como strings formatadas
func (p *PrintTable) Format(rows interface{}) string {
	t := reflect.TypeOf(rows)
	v := reflect.ValueOf(rows)

	if t.Kind() != reflect.Slice {
		fmt.Println("rows is not a slice")
		return ""
	}

	var s string
	if len(p.header) != 0 {
		s += p.mountHeader() + "\n"
	}

	for i := 0; i < v.Len(); i++ {
		s += p.toString(v.Index(i).Interface()) + "\n"
	}
	return s
}

func (p *PrintTable) mountHeader() string {
	var values = make([]interface{}, 0)
	for _, r := range p.header {
		values = append(values, r)
	}
	return fmt.Sprintf(p.layout, values...)
}

func (p *PrintTable) toString(row interface{}) string {
	vrow := reflect.ValueOf(row)
	values := make([]interface{}, vrow.NumField())
	for i := 0; i < vrow.NumField(); i++ {
		values[i] = fmt.Sprintf("%v", vrow.Field(i).Interface())
	}
	return fmt.Sprintf(p.layout, values...)
}
