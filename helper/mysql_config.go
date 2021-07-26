package helper

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/url"
	"strconv"
	"time"
)

type DBConfig struct {
	username        string
	password        string
	host            string
	port            int
	database        string
	maxIdleConn     int
	maxOpenConn     int
	connMaxLifetime time.Duration
	connMaxIdleTime time.Duration
	parameters      map[string]string
	opts            []gorm.Option
}

func NewConfig() *DBConfig {
	return &DBConfig{
		username:        "root",
		password:        "root",
		host:            "localhost",
		port:            3306,
		maxIdleConn:     10,
		maxOpenConn:     0,
		connMaxLifetime: 0,
		connMaxIdleTime: 0,
		parameters:      map[string]string{},
		opts:            []gorm.Option{},
	}
}

func (cfg *DBConfig) Username(username string) *DBConfig {
	cfg.username = username
	return cfg
}

func (cfg *DBConfig) Password(password string) *DBConfig {
	cfg.password = password
	return cfg
}

func (cfg *DBConfig) Host(host string) *DBConfig {
	cfg.host = host
	return cfg
}

func (cfg *DBConfig) Port(port int) *DBConfig {
	cfg.port = port
	return cfg
}

func (cfg *DBConfig) Database(database string) *DBConfig {
	cfg.database = database
	return cfg
}

func (cfg *DBConfig) Parameter(parameter, value string) *DBConfig {
	cfg.parameters[parameter] = value
	return cfg
}

func (cfg *DBConfig) Options(opts ...gorm.Option) *DBConfig {
	cfg.opts = append(cfg.opts, opts...)
	return cfg
}

func (cfg *DBConfig) ParseTime(val bool) *DBConfig {
	cfg.parameters["parseTime"] = strconv.FormatBool(val)
	return cfg
}

func (cfg *DBConfig) Loc(val string) *DBConfig {
	cfg.parameters["loc"] = val
	return cfg
}

func (cfg *DBConfig) Charset(val string) *DBConfig {
	cfg.parameters["charset"] = val
	return cfg
}
func (cfg *DBConfig) buildParameters() string {
	query := url.Values{}
	for k, v := range cfg.parameters {
		query.Add(k, v)
	}
	return query.Encode()
}

func (cfg *DBConfig) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", []interface{}{
		cfg.username,
		cfg.password,
		cfg.host,
		cfg.port,
		cfg.database,
		cfg.buildParameters(),
	}...)
}

func (cfg *DBConfig) DB() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(cfg.DSN()), cfg.opts...)
	if err == nil {
		sqlDB, err := db.DB()
		if err != nil {
			return nil, err
		}
		sqlDB.SetMaxIdleConns(cfg.maxIdleConn)
		sqlDB.SetMaxOpenConns(cfg.maxOpenConn)
		sqlDB.SetConnMaxLifetime(cfg.connMaxLifetime)
		sqlDB.SetConnMaxIdleTime(cfg.connMaxIdleTime)
	}
	return db, err
}
