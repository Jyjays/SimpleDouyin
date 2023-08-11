package initialize

import (
	"fmt"
	"log"

	"github.com/Jyjays/SimpleDouyin/global"
	"github.com/Jyjays/SimpleDouyin/models"

	"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var user models.User

func InitializeMysql() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalf("无法加载配置文件: %v", err)
	}

	mysqlSection := cfg.Section("mysql")
	user := mysqlSection.Key("user").String()
	password := mysqlSection.Key("password").String()
	host := mysqlSection.Key("host").String()
	port := mysqlSection.Key("port").String()
	dbname := mysqlSection.Key("dbname").String()

	// 构建 DSN 字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname)

	// 连接到数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	global.Db = db
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	database, err := global.Db.DB()
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	database.SetMaxIdleConns(10)
	database.SetMaxOpenConns(100)

	global.Db.AutoMigrate(&user)

}
