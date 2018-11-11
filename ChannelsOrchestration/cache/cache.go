package cache

import (
	lediscfg "github.com/siddontang/ledisdb/config"
	"github.com/siddontang/ledisdb/ledis"
)

var CacheDB *ledis.DB

func InitLedisDB() {
	cfg := lediscfg.NewConfigDefault()
	l, _ := ledis.Open(cfg)
	CacheDB, _ = l.Select(0)
}
