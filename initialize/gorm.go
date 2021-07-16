package initialize

import (
	"fmt"
	"yasi_audio/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func MysqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.SysUser{},
		// model.Task{},
		// model.Room_Id{},
		// model.JwtBlacklist{},
		// model.Audio_time{},
		// model.Uuid{},
	)
	fmt.Println(err)
	// if err != nil {
	// 	global.GVA_LOG.Error("register table failed", zap.Any("err", err))
	// 	os.Exit(0)
	// }
	// global.GVA_LOG.Info("register table success")
}

func GormMysql() *gorm.DB {
	dsn := "root" + ":" + "cjdsm123" + "@tcp(127.0.0.1:3306)/android?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig)); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil

		fmt.Println("error")
		return nil
	} else {
		return db
	}
}
