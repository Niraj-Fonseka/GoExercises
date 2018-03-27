package ledisDB

import (
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

type DB struct {
	*ledis.DB
}

var DBLEDIS *DB

func init() {
	DBLEDIS, _ = NewLedisDB()
}

func NewLedisDB() (*DB, error) {
	cfg := lediscfg.NewConfigDefault()
	l, _ := ledis.Open(cfg)

	database, err := l.Select(0)

	return &DB{database}, err
}
