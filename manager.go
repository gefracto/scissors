package scissors

import (
	"github.com/mohae/deepcopy"
	"errors"
)

const (
	SelectAction = iota
	InsertAction
	UpdateAction
	DeleteAction
	UpsertAction
)

type Driver interface{
	Open(name string, args []interface{}) (Table, error)
	Create(name string, args []interface{}) (Table, error)
} 

func Register(name string, driver Driver) {
	
}

var (
	ErrInvalidAction = errors.New("Error invalid action")
	ErrTblNotExists = errors.New("Table does not exists")
)

type Table interface {
	Select(key string, args []interface{})(interface{}, error)
	Update(key string, value interface{}, args []interface{}) error
	Insert(key string, value interface{}, args []interface{}) error
	Delete(key string, args []interface{}) error
	Upsert(key string, value interface{}, args []interface{}) error
	Close()
}

type Response struct {
	value	interface{}
	err		error
}

func NewResponse(value interface{}, err error) *Response {
	newValue := deepcopy.Copy(value)
	return &Response{
		value:	newValue,
		err:	err,
	}
}

type Query struct {
	tableName	string
	key			string
	args		[]interface{}
	value		interface{}
	action		int
	output		chan *Response
}

func NewQuery(action int, name string, key string, value interface{}, args []interface{}) *Query {
	newValue := deepcopy.Copy(value)
	output := make(chan *Response)
	return &Query{
		tableName: 	name,
		action:		action,
		key:		key,
		value:		newValue,
		args:		args,
		output:		output,
	}
}

type manager struct {
	input	chan *Query
	tables	map[string]Table
}

func (m *manager) Query(q *Query) (interface{}, error) {
	m.input <- q
	r := <- q.output
	close(q.output)
	return r.value, r.err
}

func (m *manager) eventLoop() {
	for {
		q := <- m.input
		table, ok := m.tables[q.tableName]
		if !ok {
			err := ErrTblNotExists
			q.output <- NewResponse(nil, err)
		}

		var value interface{}
		var err error 

		switch q.action {
			case SelectAction : value, err = table.Select(q.key, q.args)
			case InsertAction : err = table.Insert(q.key, q.value, q.args)
			case UpdateAction : err = table.Update(q.key, q.value, q.args)
			case DeleteAction : err = table.Delete(q.key, q.args)
			default : err = ErrInvalidAction
		}
		r := NewResponse(value, err)
		q.output <- r
		close(q.output)
	}
}