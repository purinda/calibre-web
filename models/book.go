package models

/* Interfaces tutorials
 * http://golangtutorials.blogspot.com.au/2011/06/interfaces-in-go.html
 *
 * SQLite3 example
 * https://github.com/mattn/go-sqlite3/blob/master/sqlite3_test/sqltest.go
 */

type Book struct {
    Id             int      `form:"-"`
    Title          string   `form:"title,text,Title:" valid:"MinSize(1);MaxSize(255)"`
    Sort           string
    Timestamp      string
    Pubdate        string
    SeriesIndex    int
    AuthorSort     string
    Isbn           string
    Lccn           string
    Path           string
    Flags          int
    Uuid           string
    HasCover       bool
    LastModified   string
}

func (b *Book) TableName() string {
    return "books"
}
