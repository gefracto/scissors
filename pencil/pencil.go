package pencil

import (
	"errors"
	"scissors"
	//"runtime"
)

var ErrNoDocs = errors.New("No document in set")

func init() {
	scissors.Register("pencil", &Driver{})
}

type Driver struct{}

func (d *Driver) Create(name string, args []interface{}) (scissors.Table, error) {
	return CreateTable(name), nil
}

func (d *Driver) Open(name string, args []interface{}) (scissors.Table, error) {
	return CreateTable(name), nil
}

type Table struct {
	Name	string
	mapping	map[string]interface{}
}

func CreateTable (name string) *Table {
	table := new(Table)
	table.Name = name
	table.mapping = make(map[string]interface{})
	return table
}

func (table *Table) Insert(key string, value interface{}, args []interface{}) error {
	table.mapping[key] = value
	return nil
}

func (table *Table) Select(key string, args []interface{}) (interface{}, error) {
	value, ok := table.mapping[key];
	if !ok {
		return nil, ErrNoDocs
	}
	return value, nil
}

func (table *Table) Update(key string, value interface{}, args []interface{}) error {
	if _, ok := table.mapping[key]; !ok {
		return ErrNoDocs
	}
	table.mapping[key] = value
	return nil
}

func (table *Table) Delete(key string, args []interface{}) error {
	delete(table.mapping, key)
	return nil
}

func (table *Table) Upsert(key string, value interface{}, args []interface{}) error {
	table.mapping[key] = value
	return nil
}

//Close() may use garbage collector to clear memory after closing table
//It can make db works slower.
func (table *Table) Close() {
	table.mapping = nil
	//runtime.GC()
}