package mysql

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"

	"github.com/jinzhu/gorm"

	// mysql 驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DefaultConfig = map[string]string{
	"charset":      "utf8",
	"parsetime":    "True",
	"loc":          "Local",
	"timeout":      "15s",
	"readTimeout":  "2s",
	"writeTimeout": "5s",
	"maxIdle":      "10",
	"maxOpenConn":  "100",
}

var (
	pool  = make(map[string]*gorm.DB)
	mutex sync.RWMutex
)

func AddDB(name string, dbConfig map[string]string) {
	conn, err := gorm.Open("mysql", buildConfigString(dbConfig))
	if err != nil {
		panic("invalid mysql configure" + err.Error())
	}
	maxIdle, _ := strconv.Atoi(dbConfig["maxIdle"])
	maxOpenConn, _ := strconv.Atoi(dbConfig["maxOpenConn"])
	conn.DB().SetMaxIdleConns(maxIdle)
	conn.DB().SetMaxOpenConns(maxOpenConn)
	//conn.DB().SetConnMaxLifetime(1 * time.Hour)
	if dbConfig["debug"] == "true" {
		// TODO 设置 log 输出为 zap
		// c.SetLogger(&logger.GormLogger{})
		conn.LogMode(true)
	}
	mutex.Lock()
	pool[name] = conn
	mutex.Unlock()
}

func GetDB(name string) *gorm.DB {
	if pool[name] == nil {
		log.Fatalf("mysql [" + name + "] gone away")
		return nil
	}
	mutex.RLock()
	db := pool[name]
	mutex.RUnlock()
	return db
}

func SetDefaultValue(keyName string, config map[string]string) {
	if v, ok := config[keyName]; !ok || v == "" {
		config[keyName] = DefaultConfig[keyName]
	}
}

func buildConfigString(config map[string]string) string {
	SetDefaultValue("charset", config)
	SetDefaultValue("parsetime", config)
	SetDefaultValue("loc", config)
	SetDefaultValue("timeout", config)
	SetDefaultValue("readTimeout", config)
	SetDefaultValue("writeTimeout", config)
	SetDefaultValue("maxIdle", config)
	SetDefaultValue("maxOpenConn", config)

	var errKeyNames []string
	for k, v := range config {
		if v == "" {
			errKeyNames = append(errKeyNames, k)
		}
	}
	if len(errKeyNames) > 0 {
		log.Fatalf("buildConfigString config err: " + strings.Join(errKeyNames, ","))
		return ""
	}
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s&timeout=%s&readTimeout=%s&writeTimeout=%s",
		config["user"],
		config["password"],
		config["host"],
		config["port"],
		config["db"],
		config["charset"],
		config["parsetime"],
		config["loc"],
		config["timeout"],
		config["readTimeout"],
		config["writeTimeout"],
	)
}
