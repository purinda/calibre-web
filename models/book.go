package models

/* Interfaces tutorials
 * http://golangtutorials.blogspot.com.au/2011/06/interfaces-in-go.html
 *
 * SQLite3 example
 * https://github.com/mattn/go-sqlite3/blob/master/sqlite3_test/sqltest.go
 */

import (
	"time"
    "fmt"
    "log"
)

type Book struct {
	Id            int       `form:"-"`
	Title         string    `form:"title,text,Title:" valid:"MinSize(1);MaxSize(255)"`
	UUID          string
	PublishedDate uint64
	HasCover      bool
	FilePath      string
}
