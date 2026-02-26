package words

// 这个文件是用来连接数据库的，使用了sqlite来连接数据库
// 主要是用来获取单词的，获取单词的描述的，以及关闭数据库连接的

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func NewDatabase() *Database {
	return &Database{}
}

type Database struct {
	Db *sql.DB
}

func (d *Database) Init(path string) error {
	var err error
	d.Db, err = sql.Open("sqlite", path)
	return err
}

func (d *Database) Connected() (bool, error) {
	if err := d.Db.Ping(); err != nil {
		return false, err
	}
	return true, nil
}

func (d *Database) Close() error {
	return d.Db.Close()
}
