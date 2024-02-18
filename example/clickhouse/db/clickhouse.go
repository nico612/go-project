package db

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"github.com/golang-migrate/migrate/v4"
	"path/filepath"
	"runtime"

	clickhouseMigrate "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"strings"
	"sync"
	"time"
)

// Options clickhouse options
type Options struct {
	ServerAddr  string
	DBName      string
	Username    string
	Password    string
	DialTimeOut time.Duration
}

func DefaultClickHouseOptions() *Options {
	return &Options{
		ServerAddr:  "localhost:9000",
		DBName:      "default",
		Username:    "default",
		Password:    "",
		DialTimeOut: 10 * time.Second,
	}
}

func NewClickConn(opts *Options) (clickhouse.Conn, error) {
	options := getClickHouseOptions(opts)
	options.MaxIdleConns = 10
	options.MaxOpenConns = 100

	conn, err := clickhouse.Open(options)

	return conn, err
}

var (
	conn driver.Conn
	once sync.Once
)

func GetClickHouseConn(opts *Options) (driver.Conn, error) {
	if opts == nil && conn == nil {
		return nil, fmt.Errorf("failed to get clickhouse Conn")
	}

	var err error
	once.Do(func() {

		conn, err = NewClickConn(opts)

		// merge database
		path := GetMigrationPath()
		fmt.Println("clickhouse migration path: ", path)

		err = RunMigration(opts, path)
	})

	if conn == nil || err != nil {
		return nil, err
	}

	return conn, nil
}

func useTLS(serverAddr string) bool {
	return strings.HasSuffix(serverAddr, "9440")
}

func getClickHouseOptions(opts *Options) *clickhouse.Options {
	options := &clickhouse.Options{
		Addr: []string{opts.ServerAddr},
		Auth: clickhouse.Auth{
			Database: opts.DBName,
			Username: opts.Username,
			Password: opts.Password,
		},
		DialTimeout: opts.DialTimeOut,
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionZSTD, // 默认使用zstd压缩
		},
	}

	if useTLS(opts.ServerAddr) {
		options.TLS = &tls.Config{}
	}

	return options
}

// 获取迁移文件路径
func GetMigrationPath() string {
	_, b, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(b), "..", "migrations")
}

// RunMigration clickhouse migration
func RunMigration(opts *Options, path string) error {
	options := getClickHouseOptions(opts)
	db := clickhouse.OpenDB(options)
	driverIns, err := clickhouseMigrate.WithInstance(db, &clickhouseMigrate.Config{
		MigrationsTableEngine: "MergeTree",
		MultiStatementEnabled: true,
	})

	// get project root path
	//migrationsPath := filepath.Join(projectpath.GetRoot(), "clickhouse", "migrations")
	m, err := migrate.NewWithDatabaseInstance(
		"file://"+path,
		opts.DBName,
		driverIns,
	)

	if err != nil {
		return fmt.Errorf("error creating clickhouse db instance for migrations: %v", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error running clickhouse migrations: %v", err)

	}

	fmt.Println("Finished clickhouse migrations for db: ", opts.DBName)

	return nil
}
