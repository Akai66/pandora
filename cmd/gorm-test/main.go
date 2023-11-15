package main

import (
	"fmt"
	"time"

	"github.com/Akai66/pandora/internal/cuslog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type result struct {
	InstanceID string `gorm:"column:id"`
	Name       string `gorm:"column:name"`
	Status     int    `gorm:"column:status"`
	NickName   string `gorm:"column:nickname"`
	Password   string `gorm:"column:password"`
	Email      string `gorm:"column:email"`
	Phone      string `gorm:"column:phone"`
}

// Options defines optsions for mysql database.
type options struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	LogLevel              int
}

// New create a new gorm db instance with the given options.
func new(opts *options) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}

func main() {
	opts := &options{
		Host:     "127.0.0.1",
		Username: "root",
		Password: "iam59!z$",
		Database: "pandora",
	}

	db, err := new(opts)
	if err != nil {
		cuslog.Fatal(err)
	}

	var records []result

	// 会根据struct中定义的gorm column标签，将查询的结果自动填充到结构体对象相应的字段上
	db.Raw("select instanceID as id,name,nickname,password,phone from user where id > ?", 0).Find(&records)
	fmt.Printf("%+v", records)
}
