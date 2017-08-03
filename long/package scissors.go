//package scissors

// import (
// 	"os"
// 	"io/ioutil"
// 	"encoding/gob"
// 	"path/filepath"
// )

// const (
// 	perm = 0666
// )

// type Table struct {
// 	file	*os.File
// 	mapping	map[string]documennt
// 	name	string
// 	decoder	*gob.Decoder
// 	encoder	*gob.Encoder
// 	mappath	string
// }

// func loadTable(tablepath, mappath string) (*Table, error) {
// 	table.mappath = mappath
// 	table := new(Table)

// 	table.mapping = make(map[string]document)

// 	mapfile, err := os.Open(table.mappath)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err := gob.NewDecoder(mapfile).Decode(table.mapping)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tablefile, err := os.OpenFile(tablepath, os.O_RDWR, perm)
// 	if err != nil {
// 		return nil, err
// 	}

// 	table.file = tablefile

// 	table.encoder = gob.NewEncoder(table.file)
// 	table.decoder = gob.NewDecoder(table.file)
// 	table.name = filepath.Base(tablepath)
	
// }


// func (table *Table) closeTable() error {
// 	table.file.Close()

// 	mapfile, err := os.OpenFile(table.mappath, os.O_TRUNC | os.O_WRONLY, perm)
// 	if err != nil {
// 		return err
// 	}

// 	err := gob.NewDecoder(mapfile).Decode(table.mapping)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (table *Table) insert(key string, val interface{}) error {
// 	start, err := table.file.Seek(0, os.SEEK_END)
// 	if err != nil {
// 		return err
// 	}

// 	if err := table.encoder.Encode(val) {
// 		return err
// 	}

// 	end, err := table.position()
// 	if err != nil {
// 		return err
// 	}

// 	table.mapping[key].start = start
// 	table.mapping[key].end = end
// }

// func (table *Table) select(key) (interface{}, error) {

// }

// func (table *Table) position() (pos int64, err error) {
// 	pos, err := table.file.Seek(0, os.SEEK_CUR)
// }

