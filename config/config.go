package config

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
}

var AppConfig *Config
var Db *sqlx.DB

func InitConfig(path string) {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	AppConfig = &Config{}

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
		return
	}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
		return
	}
	logrus.Info("---------config init success---------")
	db, err := initDBSqlx()
	if err != nil {
		log.Fatalf("Failed to initialize database, got error: %v", err)
		return
	}
	Db = db
}

func initDBSqlx() (*sqlx.DB, error) {
	db, err := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/square?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, fmt.Errorf("连接数据库失败: %v", err)
	}

	// 在连接后可以配置连接池
	db.SetMaxOpenConns(200)              // 最大打开连接数
	db.SetMaxIdleConns(1200)             // 最大空闲连接数
	db.SetConnMaxLifetime(4 * time.Hour) // 连接最大存活时间

	// 测试连接
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping失败: %v", err)
	}
	fmt.Println("成功连接到MySQL数据库!")
	return db, nil
}
