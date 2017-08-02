## Table specification

For now table package must have next instances:

### 1. Table instance: must be assertable to Table interface
```go
type Table interface {
	Select(key string, args ...interface{})(interface{}, error)
	Update(key string, value interface{}, args ...interface{}) error
	Insert(key string, value interface{}, args ...interface{}) error
	Delete(key string, args ...interface{}) error
	Upsert(key string, value interface{}, args ...interface{}) error
	Close()
}
```
In future, this interface may be expanded or changed, don't forget about it.


### 2. Table connection driver e.g. :
```go
	import "scissors"
	type Driver struct{}

	func (d *Driver) Open(tablename string, args ...interface{}) (Table, error) {
		...
	}

	func (d *Driver) Create(tablename string, args ...interface{}) (Table, error) {
		...
	}
	
	func init() {
		scissors.Register("drivername", &Driver{})
	}
	
```

If args parameter required, it must be described (their types and order) in applicated documentation.

*Additional information will be given in process of development.*



