package mybolt

import "os"

var defaultPageSize = os.Getpagesize()

type DB struct {

}

func (db *DB) String() string{
	return ""
}

