package singleton

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db   *sqlx.DB
	mu   sync.Mutex
	once sync.Once
)

// GetMysqlInsOrByLock 懒汉式，通过加锁实现
func GetMysqlInsOrByLock(dns string) (*sqlx.DB, error) {
	var err error
	if db == nil {
		mu.Lock()
		defer mu.Unlock()
		if db == nil {
			db, err = sqlx.Open("mysql", dns)
			if err != nil {
				return nil, err
			}
		}
	}
	return db, nil
}

// GetMysqlInsOrByOnce 通过once实现
func GetMysqlInsOrByOnce(dns string) (*sqlx.DB, error) {
	var err error
	if db == nil {
		once.Do(func() {
			db, err = sqlx.Open("mysql", dns)
		})
	}
	if err != nil {
		return nil, err
	}
	return db, nil
}
